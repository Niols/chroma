package a

import (
	. "github.com/Niols/chroma" // nolint
	"github.com/Niols/chroma/lexers/internal"
)

// Applescript lexer.
var Applescript = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "AppleScript",
		Aliases:   []string{"applescript"},
		Filenames: []string{"*.applescript"},
		MimeTypes: []string{},
		DotAll:    true,
	},
	applescriptRules,
))

func applescriptRules() Rules {
	return Rules{
		"root": {
			{`\s+`, Text, nil},
			{`¬\n`, LiteralStringEscape, nil},
			{`'s\s+`, Text, nil},
			{`(--|#).*?$`, Comment, nil},
			{`\(\*`, CommentMultiline, Push("comment")},
			{`[(){}!,.:]`, Punctuation, nil},
			{`(«)([^»]+)(»)`, ByGroups(Text, NameBuiltin, Text), nil},
			{`\b((?:considering|ignoring)\s*)(application responses|case|diacriticals|hyphens|numeric strings|punctuation|white space)`, ByGroups(Keyword, NameBuiltin), nil},
			{`(-|\*|\+|&|≠|>=?|<=?|=|≥|≤|/|÷|\^)`, Operator, nil},
			{`\b(and|or|is equal|equals|(is )?equal to|is not|isn't|isn't equal( to)?|is not equal( to)?|doesn't equal|does not equal|(is )?greater than|comes after|is not less than or equal( to)?|isn't less than or equal( to)?|(is )?less than|comes before|is not greater than or equal( to)?|isn't greater than or equal( to)?|(is  )?greater than or equal( to)?|is not less than|isn't less than|does not come before|doesn't come before|(is )?less than or equal( to)?|is not greater than|isn't greater than|does not come after|doesn't come after|starts? with|begins? with|ends? with|contains?|does not contain|doesn't contain|is in|is contained by|is not in|is not contained by|isn't contained by|div|mod|not|(a  )?(ref( to)?|reference to)|is|does)\b`, OperatorWord, nil},
			{`^(\s*(?:on|end)\s+)(zoomed|write to file|will zoom|will show|will select tab view item|will resize( sub views)?|will resign active|will quit|will pop up|will open|will move|will miniaturize|will hide|will finish launching|will display outline cell|will display item cell|will display cell|will display browser cell|will dismiss|will close|will become active|was miniaturized|was hidden|update toolbar item|update parameters|update menu item|shown|should zoom|should selection change|should select tab view item|should select row|should select item|should select column|should quit( after last window closed)?|should open( untitled)?|should expand item|should end editing|should collapse item|should close|should begin editing|selection changing|selection changed|selected tab view item|scroll wheel|rows changed|right mouse up|right mouse dragged|right mouse down|resized( sub views)?|resigned main|resigned key|resigned active|read from file|prepare table drop|prepare table drag|prepare outline drop|prepare outline drag|prepare drop|plugin loaded|parameters updated|panel ended|opened|open untitled|number of rows|number of items|number of browser rows|moved|mouse up|mouse moved|mouse exited|mouse entered|mouse dragged|mouse down|miniaturized|load data representation|launched|keyboard up|keyboard down|items changed|item value changed|item value|item expandable|idle|exposed|end editing|drop|drag( (entered|exited|updated))?|double clicked|document nib name|dialog ended|deminiaturized|data representation|conclude drop|column resized|column moved|column clicked|closed|clicked toolbar item|clicked|choose menu item|child of item|changed|change item value|change cell value|cell value changed|cell value|bounds changed|begin editing|became main|became key|awake from nib|alert ended|activated|action|accept table drop|accept outline drop)`, ByGroups(Keyword, NameFunction), nil},
			{`^(\s*)(in|on|script|to)(\s+)`, ByGroups(Text, Keyword, Text), nil},
			{`\b(as )(alias |application |boolean |class |constant |date |file |integer |list |number |POSIX file |real |record |reference |RGB color |script |text |unit types|(?:Unicode )?text|string)\b`, ByGroups(Keyword, NameClass), nil},
			{`\b(AppleScript|current application|false|linefeed|missing value|pi|quote|result|return|space|tab|text item delimiters|true|version)\b`, NameConstant, nil},
			{`\b(ASCII (character|number)|activate|beep|choose URL|choose application|choose color|choose file( name)?|choose folder|choose from list|choose remote application|clipboard info|close( access)?|copy|count|current date|delay|delete|display (alert|dialog)|do shell script|duplicate|exists|get eof|get volume settings|info for|launch|list (disks|folder)|load script|log|make|mount volume|new|offset|open( (for access|location))?|path to|print|quit|random number|read|round|run( script)?|say|scripting components|set (eof|the clipboard to|volume)|store script|summarize|system attribute|system info|the clipboard|time to GMT|write|quoted form)\b`, NameBuiltin, nil},
			{`\b(considering|else|error|exit|from|if|ignoring|in|repeat|tell|then|times|to|try|until|using terms from|while|with|with timeout( of)?|with transaction|by|continue|end|its?|me|my|return|of|as)\b`, Keyword, nil},
			{`\b(global|local|prop(erty)?|set|get)\b`, Keyword, nil},
			{`\b(but|put|returning|the)\b`, NameBuiltin, nil},
			{`\b(attachment|attribute run|character|day|month|paragraph|word|year)s?\b`, NameBuiltin, nil},
			{`\b(about|above|against|apart from|around|aside from|at|below|beneath|beside|between|for|given|instead of|on|onto|out of|over|since)\b`, NameBuiltin, nil},
			{`\b(accepts arrow key|action method|active|alignment|allowed identifiers|allows branch selection|allows column reordering|allows column resizing|allows column selection|allows customization|allows editing text attributes|allows empty selection|allows mixed state|allows multiple selection|allows reordering|allows undo|alpha( value)?|alternate image|alternate increment value|alternate title|animation delay|associated file name|associated object|auto completes|auto display|auto enables items|auto repeat|auto resizes( outline column)?|auto save expanded items|auto save name|auto save table columns|auto saves configuration|auto scroll|auto sizes all columns to fit|auto sizes cells|background color|bezel state|bezel style|bezeled|border rect|border type|bordered|bounds( rotation)?|box type|button returned|button type|can choose directories|can choose files|can draw|can hide|cell( (background color|size|type))?|characters|class|click count|clicked( data)? column|clicked data item|clicked( data)? row|closeable|collating|color( (mode|panel))|command key down|configuration|content(s| (size|view( margins)?))?|context|continuous|control key down|control size|control tint|control view|controller visible|coordinate system|copies( on scroll)?|corner view|current cell|current column|current( field)?  editor|current( menu)? item|current row|current tab view item|data source|default identifiers|delta (x|y|z)|destination window|directory|display mode|displayed cell|document( (edited|rect|view))?|double value|dragged column|dragged distance|dragged items|draws( cell)? background|draws grid|dynamically scrolls|echos bullets|edge|editable|edited( data)? column|edited data item|edited( data)? row|enabled|enclosing scroll view|ending page|error handling|event number|event type|excluded from windows menu|executable path|expanded|fax number|field editor|file kind|file name|file type|first responder|first visible column|flipped|floating|font( panel)?|formatter|frameworks path|frontmost|gave up|grid color|has data items|has horizontal ruler|has horizontal scroller|has parent data item|has resize indicator|has shadow|has sub menu|has vertical ruler|has vertical scroller|header cell|header view|hidden|hides when deactivated|highlights by|horizontal line scroll|horizontal page scroll|horizontal ruler view|horizontally resizable|icon image|id|identifier|ignores multiple clicks|image( (alignment|dims when disabled|frame style|scaling))?|imports graphics|increment value|indentation per level|indeterminate|index|integer value|intercell spacing|item height|key( (code|equivalent( modifier)?|window))?|knob thickness|label|last( visible)? column|leading offset|leaf|level|line scroll|loaded|localized sort|location|loop mode|main( (bunde|menu|window))?|marker follows cell|matrix mode|maximum( content)? size|maximum visible columns|menu( form representation)?|miniaturizable|miniaturized|minimized image|minimized title|minimum column width|minimum( content)? size|modal|modified|mouse down state|movie( (controller|file|rect))?|muted|name|needs display|next state|next text|number of tick marks|only tick mark values|opaque|open panel|option key down|outline table column|page scroll|pages across|pages down|palette label|pane splitter|parent data item|parent window|pasteboard|path( (names|separator))?|playing|plays every frame|plays selection only|position|preferred edge|preferred type|pressure|previous text|prompt|properties|prototype cell|pulls down|rate|released when closed|repeated|requested print time|required file type|resizable|resized column|resource path|returns records|reuses columns|rich text|roll over|row height|rulers visible|save panel|scripts path|scrollable|selectable( identifiers)?|selected cell|selected( data)? columns?|selected data items?|selected( data)? rows?|selected item identifier|selection by rect|send action on arrow key|sends action when done editing|separates columns|separator item|sequence number|services menu|shared frameworks path|shared support path|sheet|shift key down|shows alpha|shows state by|size( mode)?|smart insert delete enabled|sort case sensitivity|sort column|sort order|sort type|sorted( data rows)?|sound|source( mask)?|spell checking enabled|starting page|state|string value|sub menu|super menu|super view|tab key traverses cells|tab state|tab type|tab view|table view|tag|target( printer)?|text color|text container insert|text container origin|text returned|tick mark position|time stamp|title(d| (cell|font|height|position|rect))?|tool tip|toolbar|trailing offset|transparent|treat packages as directories|truncated labels|types|unmodified characters|update views|use sort indicator|user defaults|uses data source|uses ruler|uses threaded animation|uses title from previous column|value wraps|version|vertical( (line scroll|page scroll|ruler view))?|vertically resizable|view|visible( document rect)?|volume|width|window|windows menu|wraps|zoomable|zoomed)\b`, NameAttribute, nil},
			{`\b(action cell|alert reply|application|box|browser( cell)?|bundle|button( cell)?|cell|clip view|color well|color-panel|combo box( item)?|control|data( (cell|column|item|row|source))?|default entry|dialog reply|document|drag info|drawer|event|font(-panel)?|formatter|image( (cell|view))?|matrix|menu( item)?|item|movie( view)?|open-panel|outline view|panel|pasteboard|plugin|popup button|progress indicator|responder|save-panel|scroll view|secure text field( cell)?|slider|sound|split view|stepper|tab view( item)?|table( (column|header cell|header view|view))|text( (field( cell)?|view))?|toolbar( item)?|user-defaults|view|window)s?\b`, NameBuiltin, nil},
			{`\b(animate|append|call method|center|close drawer|close panel|display|display alert|display dialog|display panel|go|hide|highlight|increment|item for|load image|load movie|load nib|load panel|load sound|localized string|lock focus|log|open drawer|path for|pause|perform action|play|register|resume|scroll|select( all)?|show|size to fit|start|step back|step forward|stop|synchronize|unlock focus|update)\b`, NameBuiltin, nil},
			{`\b((in )?back of|(in )?front of|[0-9]+(st|nd|rd|th)|first|second|third|fourth|fifth|sixth|seventh|eighth|ninth|tenth|after|back|before|behind|every|front|index|last|middle|some|that|through|thru|where|whose)\b`, NameBuiltin, nil},
			{`"(\\\\|\\"|[^"])*"`, LiteralStringDouble, nil},
			{`\b([a-zA-Z]\w*)\b`, NameVariable, nil},
			{`[-+]?(\d+\.\d*|\d*\.\d+)(E[-+][0-9]+)?`, LiteralNumberFloat, nil},
			{`[-+]?\d+`, LiteralNumberInteger, nil},
		},
		"comment": {
			{`\(\*`, CommentMultiline, Push()},
			{`\*\)`, CommentMultiline, Pop(1)},
			{`[^*(]+`, CommentMultiline, nil},
			{`[*(]`, CommentMultiline, nil},
		},
	}
}
