package main

import (
	"crypto/tls"
	"io"
	"log"
	"strconv"
	// "math/rand"
	"net/http"
	"os"
	// "time"

	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/http2"
)

var db *sql.DB

var server = os.Getenv("server")
var port = strconv.Atoi(os.Getenv("port"))
var user = os.Getenv("user")
var password = os.Getenv("password")
var database = os.Getenv("database")

func main() {
	const (
		contactEmail = os.Getenv("contactEmail")
		domain       = os.Getenv("domain")
	)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain), //your domain here
		Cache:      autocert.DirCache("certs"),     //folder for storing certificates
		Email:      contactEmail,
	}
	r := gin.Default()

	r.Static("/iPowerAssets", "./ipower")
	r.LoadHTMLGlob("./ipower/*.html")
	r.GET("/campdoc", func(c *gin.Context) {
		c.HTML(http.StatusOK, "campdoc.html", nil)
	})
	r.GET("/ipower", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/done", func(c *gin.Context) {
		c.HTML(http.StatusOK, "done.html", nil)
	})
	r.GET("/campdocBack", func(c *gin.Context) {
		c.HTML(http.StatusOK, "campdocBack.html", nil)
	})
	r.GET("/govern", func(c *gin.Context) {
		c.HTML(http.StatusOK, "govern.html", nil)
	})
	r.GET("/allTeam", func(c *gin.Context) {
		c.HTML(http.StatusOK, "allTeam.html", nil)
	})
	// r.Static("/img", "./campdoc")
	// r.Static("/campdoc","./ipower/campdoc")
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/ipower"
		r.HandleContext(c)
	})
	runGin(r)
	staticGinGET(r)
	ginPOST(r)
	governGinGET(r)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// r.Run(":80")
	go http.ListenAndServe(":http", certManager.HTTPHandler(nil)) // 支持 http-01
	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			NextProtos:     []string{http2.NextProtoTLS, "http/1.1"},
			MinVersion:     tls.VersionTLS12,
		},
		// MaxHeaderBytes: 32 << 20,
		Handler: r,
	}
	log.Fatal(server.ListenAndServeTLS("", "")) //key and cert are comming from Let's Encrypt
}

// 執行動態GET詢問
func runGin(r *gin.Engine) {
	// 從資料庫取得該teamID的隊伍名稱
	r.GET("/api/:teamID/teamname", func(c *gin.Context) {
		id := c.Param("teamID")
		ctx := context.Background()

		// Check if database is alive.
		err := db.PingContext(ctx)
		if err != nil {
			errHandle(err, c)
			return
		}

		tsql := fmt.Sprintf("SELECT team_name FROM Team where team_id = '%s';", id)

		// Execute query
		rows, err := db.QueryContext(ctx, tsql)
		defer rows.Close()
		var team_name string
		for rows.Next() {

			err := rows.Scan(&team_name)
			if err != nil {
				errHandle(err, c)
				return
			}
		}

		c.JSON(200, gin.H{
			"teamname": team_name,
		})
	})
	// 從資料庫取得該teamID所擁有的金錢、分數和聲望值
	r.GET("/api/:teamID/team_info", func(c *gin.Context) {
		id := c.Param("teamID")
		var value int
		ctx := context.Background()

		// Check if database is alive.
		err := db.PingContext(ctx)
		if err != nil {
			errHandle(err, c)
			return
		}

		tsql := fmt.Sprintf("SELECT top 1 [value] FROM Money_record WHERE team_id = '%s' ORDER BY Money_date DESC;", id)

		// Execute query
		rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
			errHandle(err, c)
			return
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&value)
			if err != nil {
				errHandle(err, c)
				return
			}
		}
		c.JSON(200, gin.H{
			"money": value,
		})
	})
	// 從資料庫取得該teamID所擁有的各項目值以及名稱
	r.GET("/api/:teamID/item", func(c *gin.Context) {
		id := c.Param("teamID")
		i := item{}
		i.getItem(db, id, c)
		c.JSON(200, i)
	})
	// 從資料庫取得該teamID所擁有的各能力值以及名稱
	r.GET("/api/:teamID/ability", func(c *gin.Context) {
		id := c.Param("teamID")
		i := ability{}
		i.getAbility(db, id, c)
		c.JSON(200, i)
	})
	r.GET("/api/:teamID/AllSummary", func(c *gin.Context) {
		id := c.Param("teamID")
		ctx := context.Background()
		// Check if database is alive.
		err := db.PingContext(ctx)
		if err != nil {
			errHandle(err, c)
			return
		}

		tsql := `
		EXEC SumProc
		select a_sum from AllSummary WHERE team_id = @teamID
		`
		stmt, err := db.Prepare(tsql)
		defer stmt.Close()
		if err != nil {
			errHandle(err, c)
			return
		}
		row := stmt.QueryRowContext(
			ctx,
			sql.Named("teamID", id),
		)
		var abilitySum float64
		err = row.Scan(&abilitySum)
		c.JSON(200, gin.H{
			"abilitySum": abilitySum,
		})
	})
}

// 執行靜態GET詢問
func staticGinGET(r *gin.Engine) {
	// 從資料庫只取得各項目的名稱
	r.GET("/fixed/itemTitle", func(c *gin.Context) {
		i := item{}
		i.getPrice(db, c)
		c.JSON(200, i)
	})
	r.GET("/fixed/campdoc_info", func(c *gin.Context) {
		ctx := context.Background()

		// Check if database is alive.
		err := db.PingContext(ctx)
		if err != nil {
			errHandle(err, c)
			return
		}

		tsql := fmt.Sprintf("SELECT TOP (1) title,note,linkname,link,imglink FROM [dbo].[Campdoc] ORDER BY docdate DESC")

		// Execute query
		rows, err := db.QueryContext(ctx, tsql)
		defer rows.Close()
		var title, note, linkname, link, imglink string
		for rows.Next() {

			err := rows.Scan(&title, &note, &linkname, &link, &imglink)
			if err != nil {
				errHandle(err, c)
				return
			}
		}

		c.JSON(200, gin.H{
			"title":    title,
			"note":     note,
			"linkname": linkname,
			"link":     link,
			"imglink":  imglink,
		})
	})
	r.GET("/fixed/team_all", func(c *gin.Context) {
		ctx := context.Background()

		// Check if database is alive.
		err := db.PingContext(ctx)
		if err != nil {
			errHandle(err, c)
			return
		}

		tsql := fmt.Sprintf("SELECT team_name,team_id FROM Team;")

		// Execute query
		rows, err := db.QueryContext(ctx, tsql)
		defer rows.Close()
		var teamName, teamID string
		var teamNameList, teamIDList []string
		for rows.Next() {

			err := rows.Scan(&teamName, &teamID)
			if err != nil {
				errHandle(err, c)
				return
			}
			teamNameList = append(teamNameList, teamName)
			teamIDList = append(teamIDList, teamID)
		}

		c.JSON(200, gin.H{
			"teamName": teamNameList,
			"teamID":   teamIDList,
		})
	})
	r.GET("/fixed/abilityTitle", func(c *gin.Context) {
		i := ability{}
		i.getTitle(db, c)
		c.JSON(200, i)
	})
}

// 執行大會修改資料詢問
func governGinGET(r *gin.Engine) {
	r.GET("/governModify/:teamID/:id/:value", func(c *gin.Context) {
		teamID := c.Param("teamID")
		id := c.Param("id")
		value := c.Param("value")

		i, err := strconv.Atoi(value)
		if err != nil {
			errHandle(err, c)
			return
		}
		tsql := `
		-- Insert rows into table 'Govern_record' in schema '[dbo]'
		INSERT INTO [dbo].[Govern_record]
		( 
		[team_id], [id], [value]
		)
		VALUES
		( 
		@teamID, @id, @value
		)`
		ctx := context.Background()
		stmt, err := db.Prepare(tsql)
		defer stmt.Close()
		if err != nil {
			errHandle(err, c)
			return
		}
		result, err := db.ExecContext(
			ctx,
			tsql,
			sql.Named("teamID", teamID),
			sql.Named("id", id),
			sql.Named("value", i),
		)
		if err != nil {
			errHandle(err, c)
			return
		}
		r, err := result.RowsAffected()
		if err != nil {
			errHandle(err, c)
			return
		}
		c.JSON(200, gin.H{
			"returnValue": r,
		})
	})
	r.GET("/governClear", func(c *gin.Context) {
		ctx := context.Background()
		tsql := `EXEC ClearPROC @Money`
		result, err := db.ExecContext(
			ctx,
			tsql,
			sql.Named("Money", 0),
		)
		if err != nil {
			errHandle(err, c)
			return
		}
		r, err := result.RowsAffected()
		if err != nil {
			errHandle(err, c)
			return
		}
		c.JSON(200, gin.H{
			"returnValue": r,
		})
	})
}

func ginPOST(r *gin.Engine) {
	r.POST("/sendout", func(c *gin.Context) {
		ctx := context.Background()
		var sendout sendOut
		c.BindJSON(&sendout)
		tsql := `
		declare  @return int;
		EXEC @return = ItemPROC @team_id ,@item_id;
		SELECT @return as "Return";
		`
		// _, err := db.ExecContext(ctx, "ItemPROC",
		// 	sql.Named("team_id", sendout.TeamID),
		// 	sql.Named("item_id", sendout.ItemID),
		// )
		stmt, err := db.Prepare(tsql)
		defer stmt.Close()
		if err != nil {
			errHandle(err, c)
			return
		}
		row := stmt.QueryRowContext(
			ctx,
			sql.Named("team_id", sendout.TeamID),
			sql.Named("item_id", sendout.ItemID))
		var returnValue int64
		err = row.Scan(&returnValue)
		if err != nil {
			errHandle(err, c)
			return
		}

		c.JSON(200, gin.H{
			"returnValue": returnValue,
		})
	})
	r.POST("/campdocSendout", func(c *gin.Context) {
		ctx := context.Background()
		var sendcampdocout sendCampdocOut
		err := c.BindJSON(&sendcampdocout)
		if err != nil {
			errHandle(err, c)
			return
		}
		tsql := `
		INSERT INTO [dbo].[Campdoc]
		(
		[title], [note], [linkname], [link], [imglink]
		)
		VALUES
		( 
		@Title, @Note, @Linkname, @Link, @Imglink
		)
		`
		result, err := db.ExecContext(
			ctx,
			tsql,
			sql.Named("Title", sendcampdocout.Title),
			sql.Named("Note", sendcampdocout.Note),
			sql.Named("Linkname", sendcampdocout.LinkName),
			sql.Named("Link", sendcampdocout.Link),
			sql.Named("Imglink", sendcampdocout.ImgLink),
		)

		if err != nil {
			errHandle(err, c)
			return
		}
		r, err := result.RowsAffected()
		if err != nil {
			errHandle(err, c)
			return
		}
		c.JSON(200, gin.H{
			"returnValue": r,
		})
	})
}
