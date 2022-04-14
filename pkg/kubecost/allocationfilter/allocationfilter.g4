grammar allocationfilter ;

filter : '(' filter 'AND' filter ')'
       | '(' filter 'OR' filter ')'
       | COMP 'AND' COMP // precedence
       | COMP 'OR' COMP
       | 'NOT' COMP
       | COMP ;

COMP : CF2 ':' CV '=' CV '*'
     | CF2 ':' CV '=' CV
     | CF1 ':' CV '*'
     | CF1 ':' CV  ;

CF1 : 'namespace' | 'pod' ;
CF2 : 'label' | 'annotation' ;           
// COMB : 'AND' | 'OR' ;
CV: [a-zA-Z]+ ;
WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
