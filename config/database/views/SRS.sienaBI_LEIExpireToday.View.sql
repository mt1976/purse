USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBI_LEIExpireToday]    Script Date: 04/12/2021 18:16:49 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBI_LEIExpireToday]
AS
SELECT        MAX({{!SQL.SOURCE}}.BusinessDate.CurrentDate) AS Today, COUNT({{!SQL.SOURCE}}.CounterpartyExtensions.InternalId) AS Count
FROM            {{!SQL.SOURCE}}.CounterpartyExtensions LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.BusinessDate ON {{!SQL.SOURCE}}.CounterpartyExtensions.LEIExpiryDate = {{!SQL.SOURCE}}.BusinessDate.CurrentDate
WHERE        ({{!SQL.SOURCE}}.CounterpartyExtensions.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.BusinessDate.InternalDeleted IS NULL)
GROUP BY {{!SQL.SOURCE}}.BusinessDate.DealingStatus
HAVING        ({{!SQL.SOURCE}}.BusinessDate.DealingStatus = 'Permitted')
GO
