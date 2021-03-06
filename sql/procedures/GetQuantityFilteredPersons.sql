USE [Administratum]
GO
-- =============================================
-- Author:		ai
-- Create date: 2020.12.07
-- Description:	Возвращает количество потребителей с учётом наложенных фильтров
-- =============================================
CREATE PROCEDURE [dbo].[GetQuantityFilteredPersons] 
	@Login NVARCHAR(100)
	, @FamilyName NVARCHAR(100) 
	, @Name NVARCHAR(100)
	, @PatrName NVARCHAR(100)
	, @Sex BIT
WITH EXECUTE AS OWNER
AS
BEGIN
	SET NOCOUNT ON;

	SELECT COUNT(*)
	FROM dbo.Persons AS p
	WHERE	p.F_Users = @Login
				AND p.C_Family_Name LIKE CONCAT('%',ISNULL(@FamilyName, p.C_Family_Name),'%')
				AND p.C_Name LIKE CONCAT('%',ISNULL(@Name, p.C_Name),'%') 
				AND p.C_Patronymic_Name LIKE CONCAT('%',ISNULL(@PatrName,p.C_Patronymic_Name),'%') 
				AND B_Sex = ISNULL(@Sex, B_Sex)

END

