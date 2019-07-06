CREATE PROCEDURE ItemPROC (@team_id VARCHAR(100),@item_id VARCHAR(100), @ResultCount int,@ResultMoney int)
AS
begin
	Insert into Money_record (team_id,value) values (@team_id, @ResultMoney);
    Insert into Item_record (team_id,item_id,value) values (@team_id,@item_id, @ResultCount);
end