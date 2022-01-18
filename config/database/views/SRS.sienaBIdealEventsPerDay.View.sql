USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBIdealEventsPerDay]    Script Date: 04/12/2021 18:16:50 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBIdealEventsPerDay]
AS
SELECT        TOP (100) PERCENT StartInterestDate, COUNT(InternalId) AS Count
FROM            {{!SQL.SOURCE}}.DealLegsInfo
WHERE        (InternalDeleted IS NULL)
GROUP BY StartInterestDate
HAVING        (NOT (StartInterestDate IS NULL))
ORDER BY StartInterestDate
GO
