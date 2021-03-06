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

package drawing ;import (_b "github.com/unidoc/unioffice";_g "github.com/unidoc/unioffice/color";_eb "github.com/unidoc/unioffice/measurement";_e "github.com/unidoc/unioffice/schema/soo/dml";);

// LineJoin is the type of line join
type LineJoin byte ;

// SetHeight sets the height of the shape.
func (_cc ShapeProperties )SetHeight (h _eb .Distance ){_cc .ensureXfrm ();if _cc ._gc .Xfrm .Ext ==nil {_cc ._gc .Xfrm .Ext =_e .NewCT_PositiveSize2D ();};_cc ._gc .Xfrm .Ext .CyAttr =int64 (h /_eb .EMU );};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_dd *_e .CT_TextParagraphProperties ;};

// AddBreak adds a new line break to a paragraph.
func (_gb Paragraph )AddBreak (){_abd :=_e .NewEG_TextRun ();_abd .Br =_e .NewCT_TextLineBreak ();_gb ._bda .EG_TextRun =append (_gb ._bda .EG_TextRun ,_abd );};func (_d LineProperties )SetNoFill (){_d .clearFill ();_d ._c .NoFill =_e .NewCT_NoFillProperties ()};

// SetBold controls the bolding of a run.
func (_eef RunProperties )SetBold (b bool ){_eef ._fdg .BAttr =_b .Bool (b )};

// SetText sets the run's text contents.
func (_ee Run )SetText (s string ){_ee ._gdg .Br =nil ;_ee ._gdg .Fld =nil ;if _ee ._gdg .R ==nil {_ee ._gdg .R =_e .NewCT_RegularTextRun ();};_ee ._gdg .R .T =s ;};

// GetPosition gets the position of the shape in EMU.
func (_ec ShapeProperties )GetPosition ()(int64 ,int64 ){_ec .ensureXfrm ();if _ec ._gc .Xfrm .Off ==nil {_ec ._gc .Xfrm .Off =_e .NewCT_Point2D ();};return *_ec ._gc .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_ec ._gc .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;};

// Paragraph is a paragraph within a document.
type Paragraph struct{_bda *_e .CT_TextParagraph };

// X returns the inner wrapped XML type.
func (_cd Paragraph )X ()*_e .CT_TextParagraph {return _cd ._bda };func (_fdf ShapeProperties )SetSolidFill (c _g .Color ){_fdf .clearFill ();_fdf ._gc .SolidFill =_e .NewCT_SolidColorFillProperties ();_fdf ._gc .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();_fdf ._gc .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetBulletChar sets the bullet character for the paragraph.
func (_bb ParagraphProperties )SetBulletChar (c string ){if c ==""{_bb ._dd .BuChar =nil ;}else {_bb ._dd .BuChar =_e .NewCT_TextCharBullet ();_bb ._dd .BuChar .CharAttr =c ;};};

// SetFont controls the font of a run.
func (_fe RunProperties )SetFont (s string ){_fe ._fdg .Latin =_e .NewCT_TextFont ();_fe ._fdg .Latin .TypefaceAttr =s ;};func MakeShapeProperties (x *_e .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};

// X returns the inner wrapped XML type.
func (_ga LineProperties )X ()*_e .CT_LineProperties {return _ga ._c };

// RunProperties controls the run properties.
type RunProperties struct{_fdg *_e .CT_TextCharacterProperties ;};

// X returns the inner wrapped XML type.
func (_cf Run )X ()*_e .EG_TextRun {return _cf ._gdg };

// SetWidth sets the width of the shape.
func (_be ShapeProperties )SetWidth (w _eb .Distance ){_be .ensureXfrm ();if _be ._gc .Xfrm .Ext ==nil {_be ._gc .Xfrm .Ext =_e .NewCT_PositiveSize2D ();};_be ._gc .Xfrm .Ext .CxAttr =int64 (w /_eb .EMU );};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_e .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};

// X returns the inner wrapped XML type.
func (_ff ShapeProperties )X ()*_e .CT_ShapeProperties {return _ff ._gc };func (_gba ShapeProperties )LineProperties ()LineProperties {if _gba ._gc .Ln ==nil {_gba ._gc .Ln =_e .NewCT_LineProperties ();};return LineProperties {_gba ._gc .Ln };};

// SetFlipVertical controls if the shape is flipped vertically.
func (_gad ShapeProperties )SetFlipVertical (b bool ){_gad .ensureXfrm ();if !b {_gad ._gc .Xfrm .FlipVAttr =nil ;}else {_gad ._gc .Xfrm .FlipVAttr =_b .Bool (true );};};

// Run is a run within a paragraph.
type Run struct{_gdg *_e .EG_TextRun };

// Properties returns the run's properties.
func (_cg Run )Properties ()RunProperties {if _cg ._gdg .R ==nil {_cg ._gdg .R =_e .NewCT_RegularTextRun ();};if _cg ._gdg .R .RPr ==nil {_cg ._gdg .R .RPr =_e .NewCT_TextCharacterProperties ();};return RunProperties {_cg ._gdg .R .RPr };};

// SetBulletFont controls the font for the bullet character.
func (_ag ParagraphProperties )SetBulletFont (f string ){if f ==""{_ag ._dd .BuFont =nil ;}else {_ag ._dd .BuFont =_e .NewCT_TextFont ();_ag ._dd .BuFont .TypefaceAttr =f ;};};

// SetAlign controls the paragraph alignment
func (_ae ParagraphProperties )SetAlign (a _e .ST_TextAlignType ){_ae ._dd .AlgnAttr =a };

// Properties returns the paragraph properties.
func (_gd Paragraph )Properties ()ParagraphProperties {if _gd ._bda .PPr ==nil {_gd ._bda .PPr =_e .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_gd ._bda .PPr );};

// SetNumbered controls if bullets are numbered or not.
func (_bfg ParagraphProperties )SetNumbered (scheme _e .ST_TextAutonumberScheme ){if scheme ==_e .ST_TextAutonumberSchemeUnset {_bfg ._dd .BuAutoNum =nil ;}else {_bfg ._dd .BuAutoNum =_e .NewCT_TextAutonumberBullet ();_bfg ._dd .BuAutoNum .TypeAttr =scheme ;};};

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_e .CT_TextParagraph )Paragraph {return Paragraph {x }};

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_ad ShapeProperties )SetFlipHorizontal (b bool ){_ad .ensureXfrm ();if !b {_ad ._gc .Xfrm .FlipHAttr =nil ;}else {_ad ._gc .Xfrm .FlipHAttr =_b .Bool (true );};};

// SetSolidFill controls the text color of a run.
func (_fc RunProperties )SetSolidFill (c _g .Color ){_fc ._fdg .NoFill =nil ;_fc ._fdg .BlipFill =nil ;_fc ._fdg .GradFill =nil ;_fc ._fdg .GrpFill =nil ;_fc ._fdg .PattFill =nil ;_fc ._fdg .SolidFill =_e .NewCT_SolidColorFillProperties ();_fc ._fdg .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();_fc ._fdg .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetGeometry sets the shape type of the shape
func (_ddb ShapeProperties )SetGeometry (g _e .ST_ShapeType ){if _ddb ._gc .PrstGeom ==nil {_ddb ._gc .PrstGeom =_e .NewCT_PresetGeometry2D ();};_ddb ._gc .PrstGeom .PrstAttr =g ;};func (_bc LineProperties )SetSolidFill (c _g .Color ){_bc .clearFill ();_bc ._c .SolidFill =_e .NewCT_SolidColorFillProperties ();_bc ._c .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();_bc ._c .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_e .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};func (_ab LineProperties )clearFill (){_ab ._c .NoFill =nil ;_ab ._c .GradFill =nil ;_ab ._c .SolidFill =nil ;_ab ._c .PattFill =nil ;};type ShapeProperties struct{_gc *_e .CT_ShapeProperties };type LineProperties struct{_c *_e .CT_LineProperties };const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// SetLevel sets the level of indentation of a paragraph.
func (_bf ParagraphProperties )SetLevel (idx int32 ){_bf ._dd .LvlAttr =_b .Int32 (idx )};

// SetSize sets the font size of the run text
func (_cgf RunProperties )SetSize (sz _eb .Distance ){_cgf ._fdg .SzAttr =_b .Int32 (int32 (sz /_eb .HundredthPoint ));};

// X returns the inner wrapped XML type.
func (_fd ParagraphProperties )X ()*_e .CT_TextParagraphProperties {return _fd ._dd };func (_aa ShapeProperties )SetNoFill (){_aa .clearFill ();_aa ._gc .NoFill =_e .NewCT_NoFillProperties ()};

// AddRun adds a new run to a paragraph.
func (_ge Paragraph )AddRun ()Run {_af :=MakeRun (_e .NewEG_TextRun ());_ge ._bda .EG_TextRun =append (_ge ._bda .EG_TextRun ,_af .X ());return _af ;};func (_bff ShapeProperties )clearFill (){_bff ._gc .NoFill =nil ;_bff ._gc .BlipFill =nil ;_bff ._gc .GradFill =nil ;_bff ._gc .GrpFill =nil ;_bff ._gc .SolidFill =nil ;_bff ._gc .PattFill =nil ;};

// SetSize sets the width and height of the shape.
func (_de ShapeProperties )SetSize (w ,h _eb .Distance ){_de .SetWidth (w );_de .SetHeight (h )};

// SetPosition sets the position of the shape.
func (_fg ShapeProperties )SetPosition (x ,y _eb .Distance ){_fg .ensureXfrm ();if _fg ._gc .Xfrm .Off ==nil {_fg ._gc .Xfrm .Off =_e .NewCT_Point2D ();};_fg ._gc .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_b .Int64 (int64 (x /_eb .EMU ));_fg ._gc .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_b .Int64 (int64 (y /_eb .EMU ));};

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_e .EG_TextRun )Run {return Run {x }};

// SetJoin sets the line join style.
func (_bd LineProperties )SetJoin (e LineJoin ){_bd ._c .Round =nil ;_bd ._c .Miter =nil ;_bd ._c .Bevel =nil ;switch e {case LineJoinRound :_bd ._c .Round =_e .NewCT_LineJoinRound ();case LineJoinBevel :_bd ._c .Bevel =_e .NewCT_LineJoinBevel ();case LineJoinMiter :_bd ._c .Miter =_e .NewCT_LineJoinMiterProperties ();};};

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_f LineProperties )SetWidth (w _eb .Distance ){_f ._c .WAttr =_b .Int32 (int32 (w /_eb .EMU ))};func (_fdgb ShapeProperties )ensureXfrm (){if _fdgb ._gc .Xfrm ==nil {_fdgb ._gc .Xfrm =_e .NewCT_Transform2D ();};};