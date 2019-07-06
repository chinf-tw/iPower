Select ability_id,[value] From Ability_record a Where team_id = '0Jczix6Z8QnVcTSI8T6j' and Ability_date = (Select Max(b.Ability_date) From Ability_record b Where team_id = '0Jczix6Z8QnVcTSI8T6j' and a.ability_id = b.ability_id) ;

INSERT into Ability_record(team_id,ability_id,[value]) VALUES('0Jczix6Z8QnVcTSI8T6j','ability3',13);
GO
Select * From Ability_record a Where team_id = '0Jczix6Z8QnVcTSI8T6j';

Select Ability.ability_name,a.[value] FROM Ability,Ability_record a Where Ability.ability_id = a.ability_id and a.team_id = '0Jczix6Z8QnVcTSI8T6j' and a.Ability_date = (Select Max(b.Ability_date) From Ability_record b Where team_id = '0Jczix6Z8QnVcTSI8T6j' and a.ability_id = b.ability_id) ORDER BY a.ability_id;