SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		ai
-- Create date: 2020.12.05
-- Description:	Возвращает отфильтрованных Потребителей в пределах одной страницы
-- =============================================
ALTER PROCEDURE dbo.GetFilteredPaginatedCustomers -- EXEC dbo.GetFilteredPaginatedCustomers 'rode_orm', '', '', '', true, 0, 7
	@Login NVARCHAR(100)
	, @FamilyName NVARCHAR(100) 
	, @Name NVARCHAR(100)
	, @PatrName NVARCHAR(100)
	, @Sex BIT
	, @OffSetRow INT 
	, @PageSize  INT
	, @Regime INT
WITH EXECUTE AS OWNER
AS
BEGIN
	SET NOCOUNT ON;

	IF @Regime = 1
	BEGIN
		SELECT [ID], [C_Family_Name], [C_Name], [C_Patronymic_Name], [F_Users], [F_Citizenship], [B_Sex], [D_Date_Birth], [D_Date_Death]
		FROM dbo.Customers 
		WHERE	F_Users = @Login
				AND C_Family_Name LIKE CONCAT('%',ISNULL(@FamilyName, C_Family_Name),'%')
				AND C_Name LIKE CONCAT('%',ISNULL(@Name, C_Name),'%') 
				AND C_Patronymic_Name LIKE CONCAT('%',ISNULL(@PatrName,C_Patronymic_Name),'%') 
				AND B_Sex = ISNULL(@Sex, B_Sex)
		ORDER BY 1,2
		OFFSET @OffSetRow ROWS 
		FETCH NEXT @PageSize ROWS ONLY
	END
	ELSE
	BEGIN
	SELECT [ID], [C_Family_Name], [C_Name], [C_Patronymic_Name], [F_Users], [F_Citizenship], [B_Sex], [D_Date_Birth], [D_Date_Death]
		FROM dbo.Customers 
		WHERE	F_Users = @Login
				AND C_Family_Name LIKE CONCAT('%',ISNULL(@FamilyName, C_Family_Name),'%')
				AND C_Name LIKE CONCAT('%',ISNULL(@Name, C_Name),'%') 
				AND C_Patronymic_Name LIKE CONCAT('%',ISNULL(@PatrName,C_Patronymic_Name),'%') 
				AND B_Sex = ISNULL(@Sex, B_Sex)
		ORDER BY 1,2
	END
END
GO
