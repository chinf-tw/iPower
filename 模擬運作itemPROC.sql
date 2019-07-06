GO
begin
    declare @Temp_Money int;
    declare @team_id VARCHAR(100);
    
    SET @Temp_Money = 0;
    SET @team_id = '0Jczix6Z8QnVcTSI8T6j';
    select top 1 @Temp_Money = Money_record.value from Money_record where Money_record.team_id = @team_id ORDER BY money_date DESC;
    select @Temp_Money;
END


EXEC ItemPROC @team_id='HlFgMODRNerTZSUkZHiC',@item_id='item1';
SELECT * from Item_record WHERE team_id = 'HlFgMODRNerTZSUkZHiC';
SELECT * FROM Money_record WHERE team_id = 'HlFgMODRNerTZSUkZHiC';


GO
BEGIN
    declare @team_id VARCHAR(100);
    SET @team_id = 'Fj3pxFZNHaqxfHKwXTj7';

    -- EXEC ItemPROC @team_id ,@item_id='item3';
    SELECT * from Item_record WHERE team_id = @team_id;
    SELECT * FROM Money_record WHERE team_id = @team_id;
    Select Item.item_name,a.[value] FROM Item,Item_record a Where Item.item_id = a.item_id and a.team_id = @team_id and a.Item_date = (Select Max(b.Item_date) From Item_record b Where team_id = @team_id and a.item_id = b.item_id)  ORDER BY a.item_id;
END

SELECT * FROM Item;

INSERT into Money_record(team_id,[value]) VALUES('0Jczix6Z8QnVcTSI8T6j',4000);