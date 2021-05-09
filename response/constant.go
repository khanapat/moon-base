package response

const (
	SuccessCode uint64 = 200

	ErrInvalidRequestCode      uint64 = 4000
	ErrBasicAuthenticationCode uint64 = 4001

	ErrDatabaseCode uint64 = 5001
)

const (
	SuccessMessageEN           string = "Success."
	ErrInternalServerMessageEN string = "Internal server error."
	// Coin
	SuccessResetCoinMessageEN     string = "Reset Coin success."
	SuccessGetSupplyCoinMessageEN string = "Get Supply success."
	SuccessQueryHistoryMessageEN  string = "Query History success."
	SuccessBuyCoinMessageEN       string = "Buy Coin success."
	ErrBuyCoinMessageEN           string = "Cannot buy coin."
	// BasicAuthen
	ErrBasicAuthenticationMessageEN string = "Authentication failed."
	// Desc
	ErrAuthenticationDescEN string = "Unable to access data. Please check user & password"
	ErrContactAdminDescEN   string = "Please contact administrator for more information."
)

const (
	SuccessMessageTH           string = "สำเร็จ."
	ErrInternalServerMessageTH string = "มีข้อผิดพลาดภายในเซิร์ฟเวอร์."
	// Coin
	SuccessResetCoinMessageTH     string = "รีเซตเหรียญสำเร็จ."
	SuccessGetSupplyCoinMessageTH string = "แสดงจำนวนเหรียญสำเร็จ."
	SuccessQueryHistoryMessageTH  string = "แสดงประวัติการทำงานสำเร็จ."
	SuccessBuyCoinMessageTH       string = "ซื้อเหรียญสำเร็จ"
	ErrBuyCoinMessageTH           string = "ไม่สามารถซื้อเหรียญได้."
	// BasicAuthen
	ErrBasicAuthenticationMessageTH string = "ยืนยันตัวตนล้มเหลว"
	// Desc
	ErrAuthenticationDescTH string = "ไม่สามารถเข้าถึงข้อมูลได้. กรุณาตรวจสอบรหัสผู้ใช้งานใหม่อีกครั้ง"
	ErrContactAdminDescTH   string = "กรุณาติดต่อเจ้าหน้าที่ดูแลระบบเพื่อรับข้อมูลเพิ่มเติม."
)

const (
	ValidateAuthorizeTokenError string = "Invalid Token"
	ValidateFieldError          string = "Invalid Parameters"
)
