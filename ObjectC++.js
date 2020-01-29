object  : unordered collection of name and value pair : 
string to value mapping 
objects are dynamic

VERY IMPORTANT :
any value not a string , number , true , false , null or undefined is an object 

- objects mutable and manipulated by refrence 
var y = x ; if x is an object , then y stores the refrence to the object x . any manipulation to y passes to x , all all the refrences 

THREE CATEGORIES OF OBJECTS :
native objects : defined by ES5 
array , function , dates , regular expression 

host object :
defined by host , such as browser , 

user-defined : 
me !

2 CATEGOTIES OF PROPERTIES : 
own : defiend inside it 
inherited : defined inside its prototype 


THREE WAYS TO MAKE AN OBJECT : 
1. literal 

var name = {name:value , ... }

2. new keyword 
var name = new functioncallconstructor() ;

3. Object.create() : 
   every object has a second object associated withit . = prototype 
   and the object inherits properties from prototype
   Object.prototype has no protoype associated 
   
   can establish a : prototype chain 
   
   if empty object to be created then not like :
   var o2 = Object.create(null) ;
   do : 
   var o2 = Object.create(Object.prototype ) ;
 
 inherit function say we define :
 function inherit(p){
  if ( p==null) throw TypeError();
  if(Object.create)
     return Object.create(p);
  var t = typeof p;
  if(t !=="object" && t !=="function") throw TypeError();
  function f(){};
  f.prototype = p ;
  return new f();
  }
 
 
