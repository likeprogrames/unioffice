//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package elements ;import (_f "encoding/xml";_b "fmt";_a "github.com/unidoc/unioffice";);

// Validate validates the Any and its children
func (_fa *Any )Validate ()error {return _fa .ValidateWithPath ("\u0041\u006e\u0079")};func (_dbf *ElementsGroupChoice )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_bged :for {_acb ,_dbc :=d .Token ();if _dbc !=nil {return _dbc ;};switch _dc :=_acb .(type ){case _f .StartElement :switch _dc .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_cf :=NewAny ();if _gd :=d .DecodeElement (_cf ,&_dc );_gd !=nil {return _gd ;};_dbf .Any =append (_dbf .Any ,_cf );default:_a .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076",_dc .Name );if _caa :=d .Skip ();_caa !=nil {return _caa ;};};case _f .EndElement :break _bged ;case _f .CharData :};};return nil ;};func (_bcd *SimpleLiteral )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {for {_cb ,_adg :=d .Token ();if _adg !=nil {return _b .Errorf ("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s",_adg );};if _fad ,_ccg :=_cb .(_f .EndElement );_ccg &&_fad .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_ada *ElementContainer )ValidateWithPath (path string )error {for _bg ,_bge :=range _ada .Choice {if _de :=_bge .ValidateWithPath (_b .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_bg ));_de !=nil {return _de ;};};return nil ;};type ElementContainer struct{Choice []*ElementsGroupChoice ;};

// Validate validates the ElementsGroup and its children
func (_gf *ElementsGroup )Validate ()error {return _gf .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070");};type Any struct{SimpleLiteral };func (_ee *ElementContainer )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {start .Name .Local ="\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072";e .EncodeToken (start );if _ee .Choice !=nil {for _ ,_bb :=range _ee .Choice {_bb .MarshalXML (e ,_f .StartElement {});};};e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};type ElementsGroupChoice struct{Any []*Any ;};

// Validate validates the ElementsGroupChoice and its children
func (_gc *ElementsGroupChoice )Validate ()error {return _gc .ValidateWithPath ("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065");};

// Validate validates the ElementContainer and its children
func (_bf *ElementContainer )Validate ()error {return _bf .ValidateWithPath ("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072");};func NewElementsGroup ()*ElementsGroup {_db :=&ElementsGroup {};return _db };func NewSimpleLiteral ()*SimpleLiteral {_aba :=&SimpleLiteral {};return _aba };

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_abf *ElementsGroup )ValidateWithPath (path string )error {for _ba ,_deg :=range _abf .Choice {if _fac :=_deg .ValidateWithPath (_b .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_ba ));_fac !=nil {return _fac ;};};return nil ;};

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_ceb *ElementsGroupChoice )ValidateWithPath (path string )error {for _fd ,_afa :=range _ceb .Any {if _ga :=_afa .ValidateWithPath (_b .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_fd ));_ga !=nil {return _ga ;};};return nil ;};type ElementsGroup struct{Choice []*ElementsGroupChoice ;};func (_af *ElementsGroup )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_bd :for {_add ,_cd :=d .Token ();if _cd !=nil {return _cd ;};switch _ce :=_add .(type ){case _f .StartElement :switch _ce .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_eb :=NewElementsGroupChoice ();if _cdc :=d .DecodeElement (&_eb .Any ,&_ce );_cdc !=nil {return _cdc ;};_af .Choice =append (_af .Choice ,_eb );default:_a .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076",_ce .Name );if _fgb :=d .Skip ();_fgb !=nil {return _fgb ;};};case _f .EndElement :break _bd ;case _f .CharData :};};return nil ;};func NewElementContainer ()*ElementContainer {_fce :=&ElementContainer {};return _fce };type SimpleLiteral struct{};func NewElementsGroupChoice ()*ElementsGroupChoice {_ccb :=&ElementsGroupChoice {};return _ccb };

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_dde *Any )ValidateWithPath (path string )error {if _cc :=_dde .SimpleLiteral .ValidateWithPath (path );_cc !=nil {return _cc ;};return nil ;};func (_fc *Any )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_fc .SimpleLiteral =*NewSimpleLiteral ();for {_d ,_fb :=d .Token ();if _fb !=nil {return _b .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073",_fb );};if _dd ,_ca :=_d .(_f .EndElement );_ca &&_dd .Name ==start .Name {break ;};};return nil ;};func (_ab *ElementsGroup )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {if _ab .Choice !=nil {for _ ,_eeg :=range _ab .Choice {_eeg .MarshalXML (e ,_f .StartElement {});};};return nil ;};

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_gcg *SimpleLiteral )ValidateWithPath (path string )error {return nil };func (_ae *ElementContainer )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_ec :for {_da ,_fge :=d .Token ();if _fge !=nil {return _fge ;};switch _ece :=_da .(type ){case _f .StartElement :switch _ece .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_bbc :=NewElementsGroupChoice ();if _g :=d .DecodeElement (&_bbc .Any ,&_ece );_g !=nil {return _g ;};_ae .Choice =append (_ae .Choice ,_bbc );default:_a .Log ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v",_ece .Name );if _aeb :=d .Skip ();_aeb !=nil {return _aeb ;};};case _f .EndElement :break _ec ;case _f .CharData :};};return nil ;};func NewAny ()*Any {_e :=&Any {};_e .SimpleLiteral =*NewSimpleLiteral ();return _e };func (_aa *ElementsGroupChoice )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {if _aa .Any !=nil {_ac :=_f .StartElement {Name :_f .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};for _ ,_afg :=range _aa .Any {e .EncodeElement (_afg ,_ac );};};return nil ;};func (_ad *Any )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {return _ad .SimpleLiteral .MarshalXML (e ,start );};func (_dbg *SimpleLiteral )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {e .EncodeToken (start );e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};

// Validate validates the SimpleLiteral and its children
func (_dbfc *SimpleLiteral )Validate ()error {return _dbfc .ValidateWithPath ("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c");};func init (){_a .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c",NewSimpleLiteral );_a .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",NewElementContainer );_a .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0061\u006e\u0079",NewAny );_a .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070",NewElementsGroup );};