USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaAccountLadder]    Script Date: 04/12/2021 18:16:47 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaAccountLadder]
AS
SELECT        {{!SQL.SOURCE}}.CashBalance.SienaReference, {{!SQL.SOURCE}}.CashBalance.BusinessDate, {{!SQL.SOURCE}}.CashBalance.ContractNumber, {{!SQL.SOURCE}}.CashBalance.Balance, {{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Currency.AmountDp
FROM            {{!SQL.SOURCE}}.Currency LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Deals ON {{!SQL.SOURCE}}.Currency.Code = {{!SQL.SOURCE}}.Deals.DealtCcy RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.CashBalance ON {{!SQL.SOURCE}}.Deals.SienaReference = {{!SQL.SOURCE}}.CashBalance.SienaReference
WHERE        ({{!SQL.SOURCE}}.CashBalance.InternalDeleted IS NULL)
GO
