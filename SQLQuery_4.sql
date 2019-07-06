SELECT TOP (1) title,note,linkname,link,imglink
  FROM [dbo].[Campdoc] ORDER BY docdate DESC

  -- Insert rows into table 'TableName' in schema '[dbo]'
  INSERT INTO [dbo].[Campdoc]
  ( -- Columns to insert data into
   [title], [note], [linkname], [link], [imglink]
  )
  VALUES
  ( -- First row: values for the columns in the list above
   'ColumnValue1', 'ColumnValue2', 'ColumnValue3','ColumnValue1', 'ColumnValue2'
  )
  GO