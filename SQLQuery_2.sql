GO
BEGIN
    declare @team_id VARCHAR(100);
    DECLARE @return int;
    SET @team_id = 'HlFgMODRNerTZSUkZHiC';
    
    Select Ability.ability_name,a.[value] 
        FROM Ability,Ability_record a 
        Where Ability.ability_id = a.ability_id 
            and a.team_id = @team_id 
            and a.Ability_date = (
                Select Max(b.Ability_date) 
                From Ability_record b 
                Where team_id = @team_id 
                    and a.ability_id = b.ability_id
            ) 
            ORDER BY a.ability_id;
    

    EXEC @return = ItemPROC @team_id ,@item_id='item3';
    print @return;
    SELECT top 1 [value] FROM Money_record WHERE team_id = @team_id ORDER BY Money_date DESC;
    Select Item.item_name,a.[value] FROM Item,Item_record a Where Item.item_id = a.item_id and a.team_id = @team_id and a.Item_date = (Select Max(b.Item_date) From Item_record b Where team_id = @team_id and a.item_id = b.item_id)  ORDER BY a.item_id;
END

declare  @return int;
EXEC @return = ItemPROC @team_id= 'Yx3Vfch6awhIYcCbdniA' ,@item_id='item1';
SELECT @return as "Return";

GO
BEGIN
    DECLARE @x int;
    declare @team_id VARCHAR(100);
    SET @x = 10;
    SET @team_id = 'Yx3Vfch6awhIYcCbdniAiiuhui';
    SELECT @x = 
    CASE @team_id
        WHEN 'Yx3Vfch6awhIYcCbdniA' then @x*2
        else @x
    END
    SELECT @x;
END

GO
BEGIN
    DECLARE @Team_ID VARCHAR(100);
    SET @Team_ID = '';
    SELECT @Team_ID = [team_id] FROM (SELECT TOP 1 * FROM Item_record ORDER BY item_date DESC)r
    select @Team_ID;
END