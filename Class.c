class is a set of objects that inherit properties from the same prototype object 

function range(from ,to){
var r = inherit(range.methods);
r.from = from ;
r.to = to ;
return r ;

}

range.methods = { 
includes : function(x) { return this.from <= x && x <= this.to;},
foreach: function(f) {
    for ( var x = Math.ceil(this.from); x<=this.to ; x++)f(x);
},
toString: function(){
return"(" + this.from + " ... " + this.to + ")";}
};


MAP data type  : 
https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Map
 key : value pair 
 var users = new Map() 
 to add values : 
 users.set(key , value ) ;
 console.log(users) ; // displays all the value 
 users.get();
  
 2 ways of searching :
 users.keys(data );
 users.values(data) 
  
 FOROF loop : 
 syntax :: 
 for( const iterator of object){
 }
 
 for of in above case will be  : 
 for ( const [key , value ] of users.entries ){
 }
 
 
 
 making a arrayof array as a map : 
 var arrofarr = [['one':'1'],['two':'2'],['three':'3']]
 var newMap = new Map(arrofarr) ; 
 console.log(newMap);
 
 
 
