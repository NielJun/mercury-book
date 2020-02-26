(function(t){function e(e){for(var a,s,i=e[0],l=e[1],c=e[2],p=0,d=[];p<i.length;p++)s=i[p],n[s]&&d.push(n[s][0]),n[s]=0;for(a in l)Object.prototype.hasOwnProperty.call(l,a)&&(t[a]=l[a]);u&&u(e);while(d.length)d.shift()();return o.push.apply(o,c||[]),r()}function r(){for(var t,e=0;e<o.length;e++){for(var r=o[e],a=!0,i=1;i<r.length;i++){var l=r[i];0!==n[l]&&(a=!1)}a&&(o.splice(e--,1),t=s(s.s=r[0]))}return t}var a={},n={app:0},o=[];function s(e){if(a[e])return a[e].exports;var r=a[e]={i:e,l:!1,exports:{}};return t[e].call(r.exports,r,r.exports,s),r.l=!0,r.exports}s.m=t,s.c=a,s.d=function(t,e,r){s.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},s.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},s.t=function(t,e){if(1&e&&(t=s(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(s.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var a in t)s.d(r,a,function(e){return t[e]}.bind(null,a));return r},s.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return s.d(e,"a",e),e},s.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},s.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],l=i.push.bind(i);i.push=e,i=i.slice();for(var c=0;c<i.length;c++)e(i[c]);var u=l;o.push([0,"chunk-vendors"]),r()})({0:function(t,e,r){t.exports=r("56d7")},"237d":function(t,e,r){"use strict";var a=r("9615"),n=r.n(a);n.a},"56d7":function(t,e,r){"use strict";r.r(e);r("cadf"),r("551c"),r("097d");var a=r("2b0e"),n=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{attrs:{id:"app"}},[r("router-view")],1)},o=[],s={name:"App"},i=s,l=(r("8840"),r("2877")),c=Object(l["a"])(i,n,o,!1,null,null,null);c.options.__file="App.vue";var u=c.exports,p=r("8c4f"),d=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"layout-page"},[r("div",{staticClass:"main-page"},[r("router-view")],1)])},f=[],m={name:"Layout"},h=m,g=Object(l["a"])(h,d,f,!1,null,null,null);g.options.__file="index.vue";var v=g.exports,b=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"home"},[r("Tabs",{attrs:{value:"1"},on:{"on-click":t.changeTab}},t._l(t.category_list,function(e){return r("TabPane",{key:e.id,attrs:{label:e.name,id:e.id}},t._l(t.current_list,function(e){return r("Card",{key:e.question_id,attrs:{title:e.caption}},[r("p",[t._v(t._s(e.content))]),r("p"),r("p",[t._v("提问时间："+t._s(e.create_time)+" 提问者："+t._s(e.author_name))])])}))}))],1)},y=[],_={name:"home",data:function(){return{current_list:[],category_list:[]}},created:function(){this.fetchCategoryList()},methods:{fetchCategoryList:function(){var t=this;this.$http.get("/api/category/list").then(function(e){console.log(e),200==e.status?0===e.data.code?(console.log(e.data.data),t.category_list=e.data.data,t.category_list.length>0&&t.fetchQuestionList(t.category_list[0].id)):t.$Message.error(e.data.message):t.$Message.error("服务繁忙，请稍后重试")})},fetchQuestionList:function(t){console.info(t);var e=this;this.$http.get("/api/question/list?category_id="+t).then(function(t){console.log(t),200==t.status?0===t.data.code?(console.log(t.data.data),e.current_list=t.data.data):e.$Message.error(t.data.message):e.$Message.error("服务繁忙，请稍后重试")},function(t){e.$Message.error("服务繁忙，请稍后重试"),console.log(t)})},changeTab:function(t){var e=parseInt(t);this.fetchQuestionList(this.category_list[e].id)}}},x=_,$=Object(l["a"])(x,b,y,!1,null,null,null);$.options.__file="index.vue";var k=$.exports,w=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"login"},[r("div",{staticClass:"login-box"},[r("Card",{attrs:{title:"登录"}},[r("Form",[r("FormItem",{attrs:{prop:"username"}},[r("Input",{attrs:{type:"text",placeholder:"Username"},model:{value:t.form.user,callback:function(e){t.$set(t.form,"user",e)},expression:"form.user"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-person-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"password"}},[r("Input",{attrs:{type:"password",placeholder:"Password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1),r("Button",{attrs:{type:"primary",long:""},on:{click:t.submit}},[t._v("登录")]),r("router-link",{staticClass:"ivu-btn ivu-btn-default ivu-btn-long",staticStyle:{"margin-top":"10px"},attrs:{to:"/register"}},[t._v("注册")]),r("div",{staticStyle:{"margin-top":"10px","text-align":"right"}},[r("router-link",{staticClass:"ivu-btn ivu-btn-text",attrs:{to:"/"}},[t._v("返回首页")])],1)],1)],1)],1)])},I=[],C=(r("96cf"),r("1da1")),M={name:"Login",data:function(){return{form:{user:"",password:""}}},methods:{submit:function(){var t=Object(C["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,this.$http.post("/api/user/login",this.form);case 2:if(e=t.sent,console.log(e.statusCode),console.log(e),200==e.status){t.next=8;break}return this.$Message.error("登陆失败 网络错误"),t.abrupt("return");case 8:0===e.data.code?(this.$Message.success("登陆成功"),this.$router.push("/")):this.$Message.error(e.data.message);case 9:case"end":return t.stop()}},t,this)}));return function(){return t.apply(this,arguments)}}()}},O=M,j=(r("f914"),Object(l["a"])(O,w,I,!1,null,"05257cf1",null));j.options.__file="index.vue";var F=j.exports,R=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"login"},[r("div",{staticClass:"login-box"},[r("Card",{attrs:{title:"注册"}},[r("Form",[r("FormItem",{attrs:{prop:"sex"}},[r("RadioGroup",{model:{value:t.form.sex,callback:function(e){t.$set(t.form,"sex",e)},expression:"form.sex"}},[r("Radio",{attrs:{label:1}},[t._v("男")]),r("Radio",{attrs:{label:2}},[t._v("女")])],1)],1),r("FormItem",{attrs:{prop:"nickname"}},[r("Input",{attrs:{type:"text",placeholder:"昵称"},model:{value:t.form.nickname,callback:function(e){t.$set(t.form,"nickname",e)},expression:"form.nickname"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-chatbubbles-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"email"}},[r("Input",{attrs:{type:"text",placeholder:"Email"},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-mail-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"user"}},[r("Input",{attrs:{type:"text",placeholder:"Username"},model:{value:t.form.user,callback:function(e){t.$set(t.form,"user",e)},expression:"form.user"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-person-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"password"}},[r("Input",{attrs:{type:"password",placeholder:"Password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1),r("Button",{attrs:{type:"primary",long:""},on:{click:t.submit}},[t._v("注册")]),r("router-link",{staticClass:"ivu-btn ivu-btn-default ivu-btn-long",staticStyle:{"margin-top":"10px"},attrs:{to:"/login"}},[t._v("登录")]),r("div",{staticStyle:{"margin-top":"10px","text-align":"right"}},[r("router-link",{staticClass:"ivu-btn ivu-btn-text",attrs:{to:"/"}},[t._v("返回首页")])],1)],1)],1)],1)])},S=[],L={name:"Regiter",data:function(){return{form:{user:"",nickname:"",sex:1,email:"",password:""}}},methods:{submit:function(){var t=Object(C["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,this.$http.post("/api/user/register",this.form);case 2:if(e=t.sent,200==e.status){t.next=6;break}return this.$Message.error("注册失败 网络错误"),t.abrupt("return");case 6:0===e.data.code?(this.$Message.success("注册成功"),this.$router.push("/login")):this.$Message.error(e.data.message);case 7:case"end":return t.stop()}},t,this)}));return function(){return t.apply(this,arguments)}}()}},P=L,T=(r("237d"),Object(l["a"])(P,R,S,!1,null,"6ebb56d5",null));T.options.__file="index.vue";var E=T.exports,U=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",[r("Card",{attrs:{title:"提问"}},[r("Form",{attrs:{"label-position":"top",model:t.form}},[r("FormItem",{attrs:{prop:"title",label:"标题"}},[r("Input",{attrs:{type:"text",placeholder:"请输入标题"},model:{value:t.form.caption,callback:function(e){t.$set(t.form,"caption",e)},expression:"form.caption"}})],1),r("FormItem",{attrs:{prop:"category",label:"类别"}},[r("Select",{attrs:{placeholder:"请选择类别"},model:{value:t.form.category_id,callback:function(e){t.$set(t.form,"category_id",e)},expression:"form.category_id"}},t._l(t.category_list,function(e){return r("Option",{attrs:{value:e.id}},[t._v(t._s(e.name))])}))],1),r("FormItem",{attrs:{prop:"centent",label:"内容"}},[r("Input",{attrs:{type:"textarea",placeholder:"请输入内容"},model:{value:t.form.content,callback:function(e){t.$set(t.form,"content",e)},expression:"form.content"}})],1),r("Button",{attrs:{type:"primary"},on:{click:t.submit}},[t._v("提交")])],1)],1)],1)},A=[],B={name:"Ask",data:function(){return{form:{caption:"",category_id:1,content:""},category_list:[]}},created:function(){this.fetchCategoryList()},methods:{submit:function(){var t=Object(C["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,this.$http.post("api/ask/submit",this.form);case 2:if(e=t.sent,console.log(e.statusCode),console.log(e),200==e.status){t.next=8;break}return this.$Message.error("提问失败，网络错误"),t.abrupt("return");case 8:0===e.data.code?(this.$Message.success("提问成功"),this.$router.push("/")):1008===e.data.code?(this.$Message.message("请先登录"),this.$router.push("/login")):this.$Message.error(e.data.message);case 9:case"end":return t.stop()}},t,this)}));return function(){return t.apply(this,arguments)}}(),fetchCategoryList:function(){this.category_list=[{id:1,name:"技术"}];var t=this;this.$http.get("api/category/list").then(function(e){console.log(e),200==e.status?0===e.data.code?t.category_list=e.data.data:t.$Message.error(e.data.message):t.$Message.error("获取列表错误，网络有问题")})}}},Q=B,q=(r("cbc8"),Object(l["a"])(Q,U,A,!1,null,"54148a20",null));q.options.__file="index.vue";var J=q.exports;a["default"].use(p["a"]);var G=new p["a"]({routes:[{path:"/",component:v,children:[{path:"",name:"home",component:k},{path:"ask",name:"ask",component:J}]},{path:"/login",name:"login",component:F},{path:"/register",name:"register",component:E}]}),z=r("2f62");a["default"].use(z["a"]);var D=new z["a"].Store({state:{},mutations:{},actions:{}}),H=r("e069"),K=r.n(H),N=(r("dcad"),r("bc3a")),V=r.n(N),W=V.a.create({baseURL:"http://localhost:9090",timeout:1e4,headers:{"Content-Type":"application/json;charset=UTF-8"}}),X=W;a["default"].use(K.a),a["default"].prototype.$http=X,a["default"].config.productionTip=!1,new a["default"]({el:"#app",router:G,store:D,render:function(t){return t(u)}})},"6d77":function(t,e,r){},8840:function(t,e,r){"use strict";var a=r("b714"),n=r.n(a);n.a},9615:function(t,e,r){},"9a5b":function(t,e,r){},b714:function(t,e,r){},cbc8:function(t,e,r){"use strict";var a=r("9a5b"),n=r.n(a);n.a},f914:function(t,e,r){"use strict";var a=r("6d77"),n=r.n(a);n.a}});
//# sourceMappingURL=app.9a1e833d.js.map