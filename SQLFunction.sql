
create function dbo.BuyItem(@team_id VARCHAR(100),@item_id VARCHAR(100) )
returns int
as
begin
  declare @ResultMoney int; --剩餘金錢
  declare @ResultCount int; --改變數量
  declare @Temp_Money int;
  declare @Item_value int;
  declare @Item_count int;
  SET @ResultMoney = 0;
  SET @ResultCount = 0;
  SET @Temp_Money = 0;
  SET @Item_value = 0;
  SET @Item_count = 0;

  -- 依序取出 隊伍金錢, 物品金錢, 各隊最新擁有數量
  select @Temp_Money = Money_record.value from Money_record where Money_record.team_id = @team_id ORDER BY money_date DESC
  select @Item_value = Item.value from Item where Item.item_id = @item_id
  SELECT @Item_count = Item_record.value FROM Item_record 
  WHERE Item_record.team_id = @team_id and Item_record.item_id = @item_id 
  ORDER BY item_date DESC 
  
  if (@Temp_Money<@Item_value)
    return -1;
  else
    SET @ResultMoney = @Temp_Money-@Item_value;
    SET @ResultCount = @Item_count+1;
    -- 插入 改變後金錢，物品數量
    --select * from ItemPROC (@team_id,@item_id, @ResultCount, @ResultMoney)
	Exec ItemPROC @team_id,@item_id, @ResultCount, @ResultMoney
    -- Insert into Money_record (team_id,value) values (@team_id, @ResultMoney);
    -- Insert into Item_record (team_id,item_id,value) values (@team_id,@item_id, @ResultCount);
    return 0;
end


-- BEGIN
-- DECLARE @tmp DECIMAL(38,0)
--     IF (@n <= 1)
--         SELECT @tmp = 1
--  ELSE
--   SELECT @tmp = @n * dbo.fakultät(@n - 1)
--  RETURN @tmp
-- END

-- CREATE FUNCTION [dbo].[Test_Function]
--         (@ID int)
-- RETURNS @TABLE TABLE (
--         ID varchar(25),
--         Number varchar(25)
--                         )
-- AS
-- BEGIN  

--     IF @ID = 1
--         begin
--                 INSERT INTO @TABLE (ID,Number)
--                 select 'ID1','ID Number 1'
--         end
--     else
--         BEGIN
--                 INSERT INTO @TABLE (ID,Number)
--                 select 'ID2','ID Number 2'
--         end

-- RETURN
-- END