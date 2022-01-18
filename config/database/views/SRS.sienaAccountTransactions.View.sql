USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaAccountTransactions]    Script Date: 04/12/2021 18:16:49 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaAccountTransactions]
AS
SELECT        {{!SQL.SOURCE}}.DealLegsInfo.SienaReference, {{!SQL.SOURCE}}.DealLegsInfo.LegNo, {{!SQL.SOURCE}}.DealLegsInfo.MMLegNo, {{!SQL.SOURCE}}.DealLegsInfo.Narrative, {{!SQL.SOURCE}}.DealLegsInfo.Amount, {{!SQL.SOURCE}}.DealLegsInfo.StartInterestDate, {{!SQL.SOURCE}}.DealLegsInfo.EndInterestDate, 
                         {{!SQL.SOURCE}}.DealLegsInfo.Amortisation, {{!SQL.SOURCE}}.DealLegsInfo.InterestAmount, {{!SQL.SOURCE}}.DealLegsInfo.InterestAction, {{!SQL.SOURCE}}.DealLegsInfo.FixingDate, {{!SQL.SOURCE}}.DealLegsInfo.InterestCalculationDate, {{!SQL.SOURCE}}.DealLegsInfo.AmendmentAmount, 
                         {{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Currency.AmountDp
FROM            {{!SQL.SOURCE}}.Currency LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Deals ON {{!SQL.SOURCE}}.Currency.Code = {{!SQL.SOURCE}}.Deals.DealtCcy RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.DealLegsInfo ON {{!SQL.SOURCE}}.Deals.SienaReference = {{!SQL.SOURCE}}.DealLegsInfo.SienaReference
WHERE        ({{!SQL.SOURCE}}.Deals.DealType = 'Acct') AND ({{!SQL.SOURCE}}.Deals.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.DealLegsInfo.InternalDeleted IS NULL)
GO
