(function(t){function e(e){for(var a,o,i=e[0],l=e[1],c=e[2],p=0,f=[];p<i.length;p++)o=i[p],n[o]&&f.push(n[o][0]),n[o]=0;for(a in l)Object.prototype.hasOwnProperty.call(l,a)&&(t[a]=l[a]);u&&u(e);while(f.length)f.shift()();return s.push.apply(s,c||[]),r()}function r(){for(var t,e=0;e<s.length;e++){for(var r=s[e],a=!0,i=1;i<r.length;i++){var l=r[i];0!==n[l]&&(a=!1)}a&&(s.splice(e--,1),t=o(o.s=r[0]))}return t}var a={},n={app:0},s=[];function o(e){if(a[e])return a[e].exports;var r=a[e]={i:e,l:!1,exports:{}};return t[e].call(r.exports,r,r.exports,o),r.l=!0,r.exports}o.m=t,o.c=a,o.d=function(t,e,r){o.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},o.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},o.t=function(t,e){if(1&e&&(t=o(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(o.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var a in t)o.d(r,a,function(e){return t[e]}.bind(null,a));return r},o.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return o.d(e,"a",e),e},o.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},o.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],l=i.push.bind(i);i.push=e,i=i.slice();for(var c=0;c<i.length;c++)e(i[c]);var u=l;s.push([0,"chunk-vendors"]),r()})({0:function(t,e,r){t.exports=r("56d7")},"3a21":function(t,e,r){"use strict";var a=r("802c"),n=r.n(a);n.a},"56d7":function(t,e,r){"use strict";r.r(e);r("cadf"),r("551c"),r("097d");var a=r("2b0e"),n=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{attrs:{id:"app"}},[r("router-view")],1)},s=[],o={name:"App"},i=o,l=(r("8840"),r("2877")),c=Object(l["a"])(i,n,s,!1,null,null,null);c.options.__file="App.vue";var u=c.exports,p=r("8c4f"),f=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"layout-page"},[r("div",{staticClass:"main-page"},[r("router-view")],1)])},m=[],d={name:"Layout"},b=d,v=Object(l["a"])(b,f,m,!1,null,null,null);v.options.__file="index.vue";var h=v.exports,g=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"home"},[r("Tabs",{attrs:{value:"1"}},[r("TabPane",{attrs:{label:"技术",name:"1"}},t._l(t.list1,function(e){return r("Card",{key:e.id,attrs:{title:e.title}},[r("p",[t._v(t._s(e.content))]),r("p",[t._v("提问时间："+t._s(e.createAt)+" 提问者："+t._s(e.createByName))])])})),r("TabPane",{attrs:{label:"生活",name:"2"}},t._l(t.list2,function(e){return r("Card",{key:e.id,attrs:{title:e.title}},[r("p",[t._v(t._s(e.content))]),r("p",[t._v("提问时间："+t._s(e.createAt)+" 提问者："+t._s(e.createByName))])])})),r("TabPane",{attrs:{label:"八卦",name:"3"}},t._l(t.list3,function(e){return r("Card",{key:e.id,attrs:{title:e.title}},[r("p",[t._v(t._s(e.content))]),r("p",[t._v("提问时间："+t._s(e.createAt)+" 提问者："+t._s(e.createByName))])])}))],1)],1)},y=[],_={name:"home",data:function(){return{list1:[{id:1,title:"aaaa",content:"bbbbbbbbbb",createAt:"2018-10-10",createByName:"rrrrr"}],list2:[{id:1,title:"aaaa",content:"bbbbbbbbbb",craeteAt:"2018-10-10",craeteByName:"rrrrr"}],list3:[{id:1,title:"aaaa",content:"bbbbbbbbbb",craeteAt:"2018-10-10",craeteByName:"rrrrr"}]}}},x=_,k=Object(l["a"])(x,g,y,!1,null,null,null);k.options.__file="index.vue";var w=k.exports,$=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"login"},[r("div",{staticClass:"login-box"},[r("Card",{attrs:{title:"登录"}},[r("Form",[r("FormItem",{attrs:{prop:"username"}},[r("Input",{attrs:{type:"text",placeholder:"Username"},model:{value:t.form.user,callback:function(e){t.$set(t.form,"user",e)},expression:"form.user"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-person-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"password"}},[r("Input",{attrs:{type:"password",placeholder:"Password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1),r("Button",{attrs:{type:"primary",long:""},on:{click:t.submit}},[t._v("登录")]),r("router-link",{staticClass:"ivu-btn ivu-btn-default ivu-btn-long",staticStyle:{"margin-top":"10px"},attrs:{to:"/register"}},[t._v("注册")]),r("div",{staticStyle:{"margin-top":"10px","text-align":"right"}},[r("router-link",{staticClass:"ivu-btn ivu-btn-text",attrs:{to:"/"}},[t._v("返回首页")])],1)],1)],1)],1)])},I=[],C=(r("96cf"),r("1da1")),O={name:"Login",data:function(){return{form:{user:"",password:""}}},methods:{submit:function(){var t=Object(C["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return console.log("调用登录接口"),t.next=3,this.$http.post("/api/user/login",this.form);case 3:if(e=t.sent,console.log(e),200==e.status){t.next=8;break}return this.$Message.error("登陆失败，网络错误"),t.abrupt("return");case 8:0===e.data.code?(this.$Message.success("登陆成功"),this.$router.push("/")):this.$Message.error(e.data.message);case 9:case"end":return t.stop()}},t,this)}));return function(){return t.apply(this,arguments)}}()}},j=O,M=(r("968f"),Object(l["a"])(j,$,I,!1,null,"9a326230",null));M.options.__file="index.vue";var F=M.exports,P=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"login"},[r("div",{staticClass:"login-box"},[r("Card",{attrs:{title:"注册"}},[r("Form",[r("FormItem",{attrs:{prop:"sex"}},[r("RadioGroup",{model:{value:t.form.sex,callback:function(e){t.$set(t.form,"sex",e)},expression:"form.sex"}},[r("Radio",{attrs:{label:1}},[t._v("男")]),r("Radio",{attrs:{label:2}},[t._v("女")])],1)],1),r("FormItem",{attrs:{prop:"nickname"}},[r("Input",{attrs:{type:"text",placeholder:"昵称"},model:{value:t.form.nickname,callback:function(e){t.$set(t.form,"nickname",e)},expression:"form.nickname"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-chatbubbles-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"email"}},[r("Input",{attrs:{type:"text",placeholder:"Email"},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-mail-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"user"}},[r("Input",{attrs:{type:"text",placeholder:"Username"},model:{value:t.form.user,callback:function(e){t.$set(t.form,"user",e)},expression:"form.user"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-person-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"password"}},[r("Input",{attrs:{type:"password",placeholder:"Password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1),r("Button",{attrs:{type:"primary",long:""},on:{click:t.submit}},[t._v("注册")]),r("router-link",{staticClass:"ivu-btn ivu-btn-default ivu-btn-long",staticStyle:{"margin-top":"10px"},attrs:{to:"/login"}},[t._v("登录")]),r("div",{staticStyle:{"margin-top":"10px","text-align":"right"}},[r("router-link",{staticClass:"ivu-btn ivu-btn-text",attrs:{to:"/"}},[t._v("返回首页")])],1)],1)],1)],1)])},R=[],S={name:"Regiter",data:function(){return{form:{user:"",nickname:"",sex:1,email:"",password:""}}},methods:{submit:function(){var t=Object(C["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,this.$http.post("/api/user/register",this.form);case 2:if(e=t.sent,console.log(e),200==e.status){t.next=7;break}return this.$Message.error("注册失败，网络错误"),t.abrupt("return");case 7:0===e.data.code?(this.$Message.success("注册成功"),this.$router.push("/login")):this.$Message.error(e.data.message);case 8:case"end":return t.stop()}},t,this)}));return function(){return t.apply(this,arguments)}}()}},A=S,B=(r("5f60"),Object(l["a"])(A,P,R,!1,null,"700d041b",null));B.options.__file="index.vue";var T=B.exports,E=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",[r("Card",{attrs:{title:"提问"}},[r("Form",{attrs:{"label-position":"top",model:t.form}},[r("FormItem",{attrs:{prop:"title",label:"标题"}},[r("Input",{attrs:{type:"text",placeholder:"请输入标题"},model:{value:t.form.title,callback:function(e){t.$set(t.form,"title",e)},expression:"form.title"}})],1),r("FormItem",{attrs:{prop:"category",label:"类别"}},[r("Select",{attrs:{placeholder:"请选择类别"},model:{value:t.form.category,callback:function(e){t.$set(t.form,"category",e)},expression:"form.category"}},t._l(t.category_list,function(e){return r("Option",{attrs:{value:e.id}},[t._v(t._s(e.name))])}))],1),r("FormItem",{attrs:{prop:"content",label:"内容"}},[r("Input",{attrs:{type:"textarea",placeholder:"请输入内容"},model:{value:t.form.content,callback:function(e){t.$set(t.form,"content",e)},expression:"form.content"}})],1),r("Button",{attrs:{type:"primary"},on:{click:t.submit}},[t._v("提交")])],1)],1)],1)},N=[],L={name:"Ask",data:function(){return{form:{title:"",category:1,content:""}}},created:function(){this.fetchCategoryList()},methods:{submit:function(){var t=Object(C["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,this.$http.post("/api/question/submit",this.form);case 2:if(e=t.sent,console.log(e),200==e.status){t.next=7;break}return this.$Message.error("提问失败，网络错误"),t.abrupt("return");case 7:0===e.data.code?(this.$Message.success("提问成功"),this.$router.push("/")):this.$Message.error(e.data.message);case 8:case"end":return t.stop()}},t,this)}));return function(){return t.apply(this,arguments)}}(),fetchCategoryList:function(){this.$http.get("/api/category/list").then(function(t){console.log(t),200==t.status?0===t.data.code?this.category_list=t.data.data:this.$Message.error(t.data.message):this.$Message.error("服务繁忙，请稍后重试")},function(t){this.$Message.error("服务繁忙，请稍后重试"),console.log(t)})}}},U=L,J=(r("3a21"),Object(l["a"])(U,E,N,!1,null,"2bf76bf9",null));J.options.__file="index.vue";var G=J.exports;a["default"].use(p["a"]);var q=new p["a"]({routes:[{path:"/",component:h,children:[{path:"",name:"home",component:w},{path:"ask",name:"ask",component:G}]},{path:"/login",name:"login",component:F},{path:"/register",name:"register",component:T}]}),z=r("2f62");a["default"].use(z["a"]);var D=new z["a"].Store({state:{},mutations:{},actions:{}}),H=r("e069"),K=r.n(H),Q=(r("dcad"),r("bc3a")),V=r.n(Q),W=V.a.create({baseURL:"http://localhost:9090",timeout:1e4,headers:{"Content-Type":"application/json;charset=UTF-8"}}),X=W;a["default"].use(K.a),a["default"].prototype.$http=X,a["default"].config.productionTip=!1,new a["default"]({el:"#app",router:q,store:D,render:function(t){return t(u)}})},"5f60":function(t,e,r){"use strict";var a=r("ec3f"),n=r.n(a);n.a},"5f72":function(t,e,r){},"802c":function(t,e,r){},8840:function(t,e,r){"use strict";var a=r("b714"),n=r.n(a);n.a},"968f":function(t,e,r){"use strict";var a=r("5f72"),n=r.n(a);n.a},b714:function(t,e,r){},ec3f:function(t,e,r){}});
//# sourceMappingURL=app.1be38431.js.map