package client

type operation interface {
	Process(state *albionState)
}

//go:generate stringer -type=OperationType
type OperationType uint16
const (
	Unused = iota
	Ping
	Join
	CreateAccount
	Login
	SendCrashLog
	CreateCharacter
	DeleteCharacter
	SelectCharacter
	RedeemKeycode
	GetGameServerByCluster
	GetSubscriptionDetails
	GetActiveSubscription
	GetSubscriptionUrl
	GetBuyTrialDetails
	GetReferralSeasonDetails
	GetAvailableTrialKeys
	Move
	AttackStart
	CastStart
	CastCancel
	TerminateToggleSpell
	ChannelingCancel
	AttackBuildingStart
	InventoryDestroyItem
	InventoryMoveItem 					// quick equip is move
	InventorySplitStack
	ChangeCluster
	ConsoleCommand
	ChatMessage
	ReportClientError
	RegisterToObject						// building, instanciate chest
	UnRegisterFromObject				// building, instanciate chest
	CraftBuildingChangeSettings
	CraftBuildingTakeMoney
	RepairBuildingChangeSettings
	RepairBuildingTakeMoney
	ActionBuildingChangeSettings
	HarvestStart
	HarvestCancel
	TakeSilver
	ActionOnBuildingStart
	ActionOnBuildingCancel
	ItemRerollQualityStart
	ItemRerollQualityCancel
	InstallResourceStart
	InstallResourceCancel
	InstallSilver
	BuildingFillNutrition
	BuildingChangeRenovationState
	BuildingBuySkin
	BuildingClaim
	BuildingGiveup
	BuildingNutritionSilverStorageDeposit
	BuildingNutritionSilverStorageWithdraw
	BuildingNutritionSilverRewardSet
	ConstructionSiteCreate
	PlaceableItemPlace
	PlaceableItemPlaceCancel
	PlaceableObjectPickup
	FurnitureObjectUse
	FarmableHarvest
	FarmableFinishGrownItem
	FarmableDestroy
	FarmableGetProduct
	FarmableFill
	LaborerObjectPlace
	LaborerObjectPlaceCancel
	CastleGateUse
	AuctionCreateOffer
	AuctionCreateRequest
	AuctionGetOffers			// JSON
	AuctionGetRequests 		// JSON
	AuctionBuyOffer
	AuctionAbortAuction
	AuctionAbortOffer
	AuctionAbortRequest
	AuctionSellRequest
	AuctionGetFinishedAuctions
	AuctionFetchAuction
	AuctionGetMyOpenOffers
	AuctionGetMyOpenRequests
	AuctionGetMyOpenAuctions
	AuctionGetItemsAverage
	ContainerOpen
	ContainerClose
	Respawn
	Suicide
	JoinGuild
	LeaveGuild
	CreateGuild
	InviteToGuild
	DeclineGuildInvitation
	KickFromGuild
	DuellingChallengePlayer
	DuellingAcceptChallenge
	DuellingDenyChallenge
	ChangeClusterTax
	ClaimTerritory
	GiveUpTerritory
	ChangeTerritoryAccessRights
	GetMonolithInfo
	GetClaimInfo
	GetAttackInfo
	GetTerritorySeasonPoints
	GetAttackSchedule
	ScheduleAttack
	GetMatches
	GetMatchDetails
	JoinMatch
	LeaveMatch
	ChangeChatSettings
	LogoutStart
	LogoutCancel
	ClaimOrbStart
	ClaimOrbCancel
	OpenBattleVault
	CloseBattleVault
	DepositToGuildAccount
	WithdrawalFromAccount
	ChangeGuildPayUpkeepFlag
	ChangeGuildTax
	GetMyTerritories
	MorganaCommand
	GetServerInfo
	InviteMercenaryToMatch
	SubscribeToCluster
	AnswerMercenaryInvitation
	GetCharacterEquipment
	GetCharacterStats
	GetKillHistoryDetails
	LearnMasteryLevel
	ChangeAvatar
	PromotePlayer
	DemotePlayer
	GetRankings
	GetRank
	GetGvgSeasonRankings
	GetGvgSeasonRank
	GetGvgSeasonHistoryRankings
	KickFromGvGMatch
	InviteToPlayerTrade
	PlayerTradeCancel
	PlayerTradeInvitationAccept
	PlayerTradeAddItem
	PlayerTradeRemoveItem
	PlayerTradeAcceptTrade
	PlayerTradeSetSilverOrGold
	SendMiniMapPing
	Stuck
	BuyRealEstate
	ClaimRealEstate
	GiveUpRealEstate
	ChangeRealEstateOutline
	GetMailInfos
	ReadMail
	SendNewMail
	DeleteMail
	ClaimAttachmentFromMail
	UpdateLfgInfo
	GetLfgInfos
	GetMyGuildLfgInfo
	GetLfgDescriptionText
	LfgApplyToGuild
	AnswerLfgGuildApplication
	GetClusterInfo
	RegisterChatPeer
	SendChatMessage
	JoinChatChannel
	LeaveChatChannel
	SendWhisperMessage
	Say
	PlayEmote
	StopEmote
	GetClusterMapInfo
	AccessRightsChangeSettings
	Mount
	MountCancel
	BuyJourney
	SetSaleStatusForEstate
	ResolveGuildOrPlayerName
	GetRespawnInfos
	MakeHome
	LeaveHome
	ResurrectionReply
	AllianceCreate
	AllianceDisband
	AllianceGetMemberInfos
	AllianceInvite
	AllianceAnswerInvitation
	AllianceCancelInvitation
	AllianceKickGuild
	AllianceLeave
	AllianceChangeGoldPaymentFlag
	AllianceGetDetailInfo
	GetIslandInfos
	AbandonMyIsland
	BuyMyIsland
	BuyGuildIsland
	AbandonGuildIsland
	UpgradeMyIsland
	UpgradeGuildIsland
	MonolithBankOpen
	MonolithBankClose
	TerritoryFillNutrition
	TeleportBack
	PartyInvitePlayer
	PartyAnswerInvitation
	PartyLeave
	PartyKickPlayer
	PartyMakeLeader
	PartyChangeLootSetting
	GetGuildMOTD
	SetGuildMOTD
	ExitEnterStart
	ExitEnterCancel
	AgentRequest
	GoldMarketGetBuyOffer
	GoldMarketGetBuyOfferFromSilver
	GoldMarketGetSellOffer
	GoldMarketGetSellOfferFromSilver
	GoldMarketBuyGold
	GoldMarketSellGold
	GoldMarketCreateSellOrder
	GoldMarketCreateBuyOrder
	GoldMarketGetInfos
	GoldMarketCancelOrder
	GoldMarketGetAverageInfo
	SiegeCampClaimStart
	SiegeCampClaimCancel
	SiegeCampBankOpen
	SiegeCampBankClose
	ChangeUseCraftingFocus
	TreasureChestUsingStart
	TreasureChestUsingCancel
	LaborerStartJob
	LaborerTakeJobLoot
	LaborerDismiss
	LaborerMove
	LaborerBuyItem
	LaborerUpgrade
	BuyPremium
	BuyTrial
	RealEstateGetAuctionData
	RealEstateBidOnAuction
	GetSiegeCampCooldown
	FriendInvite
	FriendAnswerInvitation
	FriendCancelnvitation
	FriendRemove
	InventoryStack
	InventorySort
	EquipmentItemChangeSpell
	ExpeditionRegister
	ExpeditionRegisterCancel
	JoinExpedition
	DeclineExpeditionInvitation
	VoteStart
	VoteDoVote
	RatingDoRate
	EnteringExpeditionStart
	EnteringExpeditionCancel
	ActivateExpeditionCheckPoint
	ArenaRegister
	ArenaRegisterCancel
	ArenaLeave
	JoinArenaMatch
	DeclineArenaInvitation
	EnteringArenaStart
	EnteringArenaCancel
	UpdateCharacterStatement
	BoostFarmable
	GetStrikeHistory
	UseFunction
	UsePortalEntrance
	QueryPortalBinding
	ClaimPaymentTransaction
	ChangeUseAutoLP
	ClientPerformanceStats
	ExtendedHardwareStats
	TerritoryClaimStart
	TerritoryClaimCancel
	RequestAppStoreProducts
	VerifyProductPurchase
	QueryGuildPlayerStats
	TrackAchievements
	DepositItemToGuildToken
	WithdrawalItemFromGuildToken
)

