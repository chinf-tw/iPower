

INSERT into Item_record(team_id,item_id,[value]) VALUES('0Jczix6Z8QnVcTSI8T6j','item1',12);
GO
Select * From Item_record a Where team_id = '0Jczix6Z8QnVcTSI8T6j' ;

Select item_id,[value] From Item_record a  Where team_id = '0Jczix6Z8QnVcTSI8T6j' and Item_date = (Select Max(b.Item_date) From Item_record b Where team_id = '0Jczix6Z8QnVcTSI8T6j' and a.item_id = b.item_id) ORDER BY a.item_id;



GO
BEGIN
    declare @team_id VARCHAR(100);
    SET @team_id = 'Fj3pxFZNHaqxfHKwXTj7';
    Select Item.item_name,a.[value] FROM Item,Item_record a Where Item.item_id = a.item_id and a.team_id = @team_id and a.Item_date = (Select Max(b.Item_date) From Item_record b Where team_id = @team_id and a.item_id = b.item_id)  ORDER BY a.item_id;
END