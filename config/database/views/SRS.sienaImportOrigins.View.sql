USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaImportOrigins]    Script Date: 04/12/2021 18:16:56 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaImportOrigins]
AS
SELECT DISTINCT KeyOriginID
FROM            {{!SQL.SOURCE}}.CounterpartyImportID
GO
