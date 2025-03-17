package trinity

type TokenType int

const (
	//Types
	INT        = iota //done
	FLOAT             //done
	STRING            //done
	FUNC_TYPE         // done
	NULL              // done
	BOOL              // done
	ERROR_TYPE        // done

	//Arithmetic
	PLUS     // done
	MINUS    // done
	MULTIPLY // done
	DIVIDE   //done

	//Comparison
	GREATER    //done
	GREATER_EQ //done
	LESS       //done
	LESS_EQ    //done
	EQ_EQ      //done
	NOT_EQ     //done

	//Logical
	NOT //done
	AND //done
	OR  //done

	//Single Chars
	EQUAL        //done
	SEMI_COLON   //done
	COMMA        //done
	PAREN_LEFT   //done
	PAREN_RIGHT  //done
	BRACE_LEFT   //done
	BRACE_RIGHT  //done
	DUB_QUOTES   //done
	SINGLE_QUOTE // done
	COLON        //done

	//Keywords
	MUT        //done
	CONST      //done
	IF         //done
	ELSE       //done
	ELSE_IF    //done
	FOR        //done
	INF        //done
	RANGE      // done
	PRINT      //done
	ERROR_WORD // dunno about why i added this
	MAIN_FUNC  //done

	//End of Line
	EOF //done

	//Function Related
	FUNC_RETURN // done
	FUNC_WORD   //done
)
