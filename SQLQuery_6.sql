Select Ability.ability_name,a.[value] FROM Ability,Ability_record a Where Ability.ability_id = a.ability_id and a.team_id = 'HlFgMODRNerTZSUkZHiC' and a.Ability_date = (Select Max(b.Ability_date) From Ability_record b Where team_id = 'HlFgMODRNerTZSUkZHiC' and a.ability_id = b.ability_id) ORDER BY a.ability_id;
SELECT top 1 [value] FROM Money_record WHERE team_id = 'HlFgMODRNerTZSUkZHiC' ORDER BY Money_date DESC;
GO
-- Insert rows into table 'TableName' in schema '[dbo]'
INSERT INTO [dbo].[Govern_record]
( -- Columns to insert data into
 [team_id], [id], [value]
)
VALUES
( -- First row: values for the columns in the list above
 'HlFgMODRNerTZSUkZHiC', 'ability4', 2
)
INSERT INTO [dbo].[Govern_record]
( -- Columns to insert data into
 [team_id], [id], [value]
)
VALUES
( -- First row: values for the columns in the list above
 'HlFgMODRNerTZSUkZHiC', 'ability4', 2
)
INSERT INTO [dbo].[Govern_record]
( -- Columns to insert data into
 [team_id], [id], [value]
)
VALUES
( -- First row: values for the columns in the list above
 'HlFgMODRNerTZSUkZHiC', 'ability4', 2
)
GO
SELECT * FROM Govern_record;
GO
SELECT top 1 [value] FROM Money_record WHERE team_id = 'HlFgMODRNerTZSUkZHiC' ORDER BY Money_date DESC;
GO
Select Ability.ability_name,a.[value] FROM Ability,Ability_record a Where Ability.ability_id = a.ability_id and a.team_id = 'HlFgMODRNerTZSUkZHiC' and a.Ability_date = (Select Max(b.Ability_date) From Ability_record b Where team_id = 'HlFgMODRNerTZSUkZHiC' and a.ability_id = b.ability_id) ORDER BY a.ability_id;
GO