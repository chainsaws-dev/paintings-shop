!function(){function n(n,e){if(!(n instanceof e))throw new TypeError("Cannot call a class as a function")}function e(n,e){for(var t=0;t<e.length;t++){var o=e[t];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(n,o.key,o)}}(window.webpackJsonp=window.webpackJsonp||[]).push([[12],{xYuj:function(t,o,i){"use strict";i.r(o),i.d(o,"TotpModule",function(){return S});var r=i("ofXK"),a=i("tyNb"),c=i("PCNd"),s=i("3Pt+"),u=i("1kSV"),b=i("fXoL"),d=i("qXBG");function p(n,e){if(1&n){var t=b.Ub();b.Tb(0,"ngb-alert",17),b.dc("close",function(){return b.tc(t),b.fc(3).ShowMessage=!1}),b.Cc(1),b.Sb()}if(2&n){var o=b.fc(3);b.kc("type",o.MessageType),b.Bb(1),b.Dc(o.ResponseFromBackend.Error.Message)}}function g(n,e){if(1&n&&(b.Tb(0,"div",15),b.Ac(1,p,2,2,"ngb-alert",16),b.Sb()),2&n){var t=b.fc(2);b.Bb(1),b.kc("ngIf",t.ShowMessage)}}function f(n,e){if(1&n){var t=b.Ub();b.Tb(0,"div",4),b.Tb(1,"div",5),b.Tb(2,"h3"),b.Cc(3,"Second factor"),b.Sb(),b.Sb(),b.Tb(4,"div",6),b.Tb(5,"form",7,8),b.dc("ngSubmit",function(){b.tc(t);var n=b.sc(6);return b.fc().OnSubmitForm(n)}),b.Tb(7,"div",9),b.Pb(8,"input",10),b.Sb(),b.Tb(9,"div",11),b.Tb(10,"div",12),b.Tb(11,"button",13),b.Cc(12,"Check"),b.Sb(),b.Sb(),b.Sb(),b.Sb(),b.Sb(),b.Ac(13,g,2,1,"div",14),b.Sb()}if(2&n){var o=b.sc(6),i=b.fc();b.Bb(11),b.kc("disabled",o.invalid),b.Bb(2),b.kc("ngIf",i.ShowMessage)}}function l(n,e){1&n&&(b.Tb(0,"div",18),b.Tb(1,"span",19),b.Cc(2,"Loading..."),b.Sb(),b.Sb())}var h,m,v=[{path:"",component:(h=function(){function t(e,o){n(this,t),this.authservice=e,this.router=o,this.LoginMode=!0,this.IsLoading=!1}var o,i,r;return o=t,(i=[{key:"ngOnDestroy",value:function(){this.SfResultSub.unsubscribe(),this.SfErrSub.unsubscribe()}},{key:"ngOnInit",value:function(){var n=this;this.authservice.CheckRegistered()&&this.Redirect(),this.SfErrSub=this.authservice.SfErrorSub.subscribe(function(e){if(n.ShowMessage=!0,n.ResponseFromBackend=e,setTimeout(function(){return n.ShowMessage=!1},5e3),e)switch(e.Error.Code){case 200:n.MessageType="success";break;default:n.MessageType="danger"}n.IsLoading=!1}),this.SfResultSub=this.authservice.SfResultSub.subscribe(function(e){n.IsLoading=!1,e&&n.Redirect()})}},{key:"OnSubmitForm",value:function(n){this.IsLoading=!0,this.authservice.SecondFactorCheck(n.value.passkey),n.reset()}},{key:"Redirect",value:function(){this.router.navigate(["/recipes"])}}])&&e(o.prototype,i),r&&e(o,r),t}(),h.\u0275fac=function(n){return new(n||h)(b.Ob(d.a),b.Ob(a.c))},h.\u0275cmp=b.Ib({type:h,selectors:[["app-totp"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","form-group"],["type","text","name","passkey","inputmode","numeric","pattern","[0-9]*","autocomplete","one-time-code","placeholder","Passkey","ngModel","","required","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars",1,"form-control"],[1,"form-group","float-right"],[1,"input-group"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],["style","padding: 3px",4,"ngIf"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(n,e){1&n&&(b.Tb(0,"div",0),b.Ac(1,f,14,2,"div",1),b.Tb(2,"div",2),b.Ac(3,l,3,0,"div",3),b.Sb(),b.Sb()),2&n&&(b.Bb(1),b.kc("ngIf",!e.IsLoading),b.Bb(2),b.kc("ngIf",e.IsLoading))},directives:[r.l,s.s,s.j,s.k,s.c,s.o,s.i,s.l,s.p,s.f,u.k,u.a],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),h)}],S=((m=function e(){n(this,e)}).\u0275mod=b.Mb({type:m}),m.\u0275inj=b.Lb({factory:function(n){return new(n||m)},imports:[[r.c,s.e,a.g.forChild(v),c.a,u.b,u.l]]}),m)}}])}();