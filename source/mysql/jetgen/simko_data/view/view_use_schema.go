//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

// UseSchema sets a new schema name for all generated view SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	AProsesServer = AProsesServer.FromSchema(schema)
	MyanggotaViewLembaga = MyanggotaViewLembaga.FromSchema(schema)
	MyanggotaViewRek = MyanggotaViewRek.FromSchema(schema)
	PengajuanAktifView = PengajuanAktifView.FromSchema(schema)
	ViewDataNasabah = ViewDataNasabah.FromSchema(schema)
	ViewProsesServer = ViewProsesServer.FromSchema(schema)
	WhatsappGetToken = WhatsappGetToken.FromSchema(schema)
	WhatsappGetTokenLembaga = WhatsappGetTokenLembaga.FromSchema(schema)
}
