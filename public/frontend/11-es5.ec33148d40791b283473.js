!function(){function e(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function t(e,t){for(var n=0;n<t.length;n++){var i=t[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(e,i.key,i)}}function n(e,n,i){return n&&t(e.prototype,n),i&&t(e,i),e}(window.webpackJsonp=window.webpackJsonp||[]).push([[11],{CXQP:function(t,i,s){"use strict";s.r(i),s.d(i,"ShoppingListModule",function(){return w});var r=s("ofXK"),o=s("AytR"),c=s("fXoL"),a=s("ozzT"),u=s("tyNb"),l=s("GXvH"),b=s("/1gQ"),d=s("3Pt+"),g=["f"];function p(e,t){if(1&e){var n=c.Ub();c.Tb(0,"button",14),c.dc("click",function(){return c.tc(n),c.fc().DeleteSelectedItem()}),c.Cc(1,"Delete"),c.Sb()}}function h(e,t){if(1&e){var n=c.Ub();c.Tb(0,"button",15),c.dc("click",function(){return c.tc(n),c.fc().ClearAllItems()}),c.Cc(1,"Clear"),c.Sb()}}var S,f=((S=function(){function t(n,i){e(this,t),this.ShopListServ=n,this.DataServ=i,this.editmode=!1}return n(t,[{key:"ngOnInit",value:function(){var e=this;this.ingselected=this.ShopListServ.IngredientSelected.subscribe(function(t){e.selectedingredient=t,e.editmode=!0,e.slEditForm.setValue({name:e.selectedingredient.Name,amount:e.selectedingredient.Amount})}),this.IngAdd=this.ShopListServ.IngredientAdded.subscribe(function(t){e.DataServ.SaveShoppingList(t)}),this.IngUpd=this.ShopListServ.IngredientUpdated.subscribe(function(t){e.DataServ.SaveShoppingList(t)}),this.IngDel=this.ShopListServ.IngredientDeleted.subscribe(function(t){e.DataServ.DeleteShoppingList(t)}),this.IngCle=this.ShopListServ.IngredientClear.subscribe(function(){e.DataServ.DeleteAllShoppingList()})}},{key:"ngOnDestroy",value:function(){this.ingselected.unsubscribe(),this.IngAdd.unsubscribe(),this.IngUpd.unsubscribe(),this.IngDel.unsubscribe(),this.IngCle.unsubscribe()}},{key:"AddNewItem",value:function(e){if(e.valid){var t=e.value;this.ShopListServ.AddNewItem(new b.b(t.name,parseInt(t.amount,10)),!1)}}},{key:"UpdateItem",value:function(e){if(e.valid){var t=e.value;this.ShopListServ.UpdateSelectedItem(new b.b(t.name,parseInt(t.amount,10))),this.editmode=!1,this.slEditForm.reset()}}},{key:"DeleteSelectedItem",value:function(){this.ShopListServ.DeleteSelectedItem()}},{key:"ClearAllItems",value:function(){this.ShopListServ.ClearAll()}}]),t}()).\u0275fac=function(e){return new(e||S)(c.Ob(a.a),c.Ob(l.a))},S.\u0275cmp=c.Ib({type:S,selectors:[["app-shopping-edit"]],viewQuery:function(e,t){var n;1&e&&c.Gc(g,!0),2&e&&c.rc(n=c.ec())&&(t.slEditForm=n.first)},decls:18,vars:4,consts:[[1,"row"],[1,"col"],[3,"ngSubmit"],["f","ngForm"],[1,"input-group","mt-3"],[1,"col-sm-9","form-group"],["type","text","id","name","placeholder","Name","name","name","ngModel","","required","",1,"form-control"],[1,"col","form-group"],["type","number","id","amount","placeholder","Amount","name","amount","ngModel","","required","","pattern","^[1-9]+[0-9]*$",1,"form-control"],[1,"input-group"],[1,"input-group-prepend"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["class","btn btn-outline-danger","type","button",3,"click",4,"ngIf"],["class","btn btn-outline-secondary","type","button",3,"click",4,"ngIf"],["type","button",1,"btn","btn-outline-danger",3,"click"],["type","button",1,"btn","btn-outline-secondary",3,"click"]],template:function(e,t){if(1&e){var n=c.Ub();c.Tb(0,"div",0),c.Tb(1,"div",1),c.Tb(2,"form",2,3),c.dc("ngSubmit",function(){c.tc(n);var e=c.sc(3);return t.editmode?t.UpdateItem(e):t.AddNewItem(e)}),c.Tb(4,"div",0),c.Tb(5,"div",4),c.Tb(6,"div",5),c.Pb(7,"input",6),c.Sb(),c.Tb(8,"div",7),c.Pb(9,"input",8),c.Sb(),c.Sb(),c.Sb(),c.Tb(10,"div",0),c.Tb(11,"div",1),c.Tb(12,"div",9),c.Tb(13,"div",10),c.Tb(14,"button",11),c.Cc(15),c.Sb(),c.Ac(16,p,2,0,"button",12),c.Ac(17,h,2,0,"button",13),c.Sb(),c.Sb(),c.Sb(),c.Sb(),c.Sb(),c.Sb(),c.Sb()}if(2&e){var i=c.sc(3);c.Bb(14),c.kc("disabled",i.invalid),c.Bb(1),c.Dc(t.editmode?"Update":"Add"),c.Bb(1),c.kc("ngIf",t.ShopListServ.CurrentSelectedItem),c.Bb(1),c.kc("ngIf",0!==t.ShopListServ.GetIngredientsLength())}},directives:[d.s,d.j,d.k,d.c,d.i,d.l,d.p,d.n,d.o,r.l],styles:["input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),S),v=s("1kSV");function m(e,t){if(1&e){var n=c.Ub();c.Tb(0,"ngb-alert",9),c.dc("close",function(){return c.tc(n),c.fc(2).ShowMessage=!1}),c.Cc(1),c.Sb()}if(2&e){var i=c.fc(2);c.kc("type",i.MessageType),c.Bb(1),c.Dc(i.ResponseFromBackend.Error.Message)}}function I(e,t){if(1&e){var n=c.Ub();c.Tb(0,"a",10),c.dc("click",function(){c.tc(n);var e=t.$implicit;return c.fc(2).ShopListServ.SelectItemShopList(e)}),c.Cc(1),c.Tb(2,"span",11),c.Cc(3),c.Sb(),c.Sb()}if(2&e){var i=t.$implicit,s=c.fc(2);c.kc("ngClass",s.ShopListServ.IsCurrentSelected(i)?"active":""),c.Bb(1),c.Ec("",i.Name," "),c.Bb(2),c.Dc(i.Amount)}}function L(e,t){if(1&e){var n=c.Ub();c.Tb(0,"ngb-pagination",12),c.dc("pageChange",function(e){return c.tc(n),c.fc(2).OnPageChanged(e)}),c.Sb()}if(2&e){var i=c.fc(2);c.kc("collectionSize",i.slcollectionSize)("pageSize",i.slPageSize)("page",i.slCurrentPage)("rotate",!0)("boundaryLinks",!0)("maxSize",11)}}function C(e,t){if(1&e&&(c.Tb(0,"div",3),c.Tb(1,"div",4),c.Pb(2,"app-shopping-edit"),c.Pb(3,"hr"),c.Ac(4,m,2,2,"ngb-alert",5),c.Tb(5,"ul",6),c.Ac(6,I,4,3,"a",7),c.Sb(),c.Ac(7,L,1,6,"ngb-pagination",8),c.Sb(),c.Sb()),2&e){var n=c.fc();c.Bb(4),c.kc("ngIf",n.ShowMessage),c.Bb(2),c.kc("ngForOf",n.ingredients),c.Bb(1),c.kc("ngIf",n.slcollectionSize>n.slPageSize)}}function y(e,t){1&e&&(c.Tb(0,"div",13),c.Tb(1,"span",14),c.Cc(2,"Loading..."),c.Sb(),c.Sb())}var k,T,P=[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:(k=function(){function t(n,i,s,r){e(this,t),this.ShopListServ=n,this.activeroute=i,this.DataServ=s,this.router=r}return n(t,[{key:"ngOnDestroy",value:function(){this.IngChanged.unsubscribe(),this.PageChanged.unsubscribe(),this.FetchOnInint.unsubscribe(),this.DataLoading.unsubscribe(),this.RecivedErrorSub.unsubscribe(),this.WatchIngAdd.unsubscribe(),this.WatchIngDel.unsubscribe(),this.WatchIngCle.unsubscribe()}},{key:"ngOnInit",value:function(){var e=this;this.slPageSize=o.a.ShoppingListPageSize,this.RecivedErrorSub=this.DataServ.RecivedError.subscribe(function(t){if(e.ShowMessage=!0,e.ResponseFromBackend=t,setTimeout(function(){return e.ShowMessage=!1},5e3),t)switch(t.Error.Code){case 200:e.MessageType="success";break;default:e.MessageType="danger"}}),this.IngChanged=this.ShopListServ.IngredientChanged.subscribe(function(t){e.ingredients=t}),this.PageChanged=this.activeroute.params.subscribe(function(t){e.slCurrentPage=+t.pn}),this.DataLoading=this.DataServ.LoadingData.subscribe(function(t){e.IsLoading=t}),this.FetchOnInint=this.DataServ.FetchShoppingList(this.slCurrentPage,o.a.ShoppingListPageSize).subscribe(function(t){e.ingredients=e.ShopListServ.GetIngredients(),e.slcollectionSize=e.ShopListServ.Total},function(t){e.ingredients=[]}),this.WatchIngAdd=this.ShopListServ.IngredientAdded.subscribe(function(t){e.slcollectionSize+=1,e.ingredients=e.ShopListServ.GetIngredients()}),this.WatchIngDel=this.ShopListServ.IngredientDeleted.subscribe(function(t){e.slcollectionSize-=1,e.ingredients=e.ShopListServ.GetIngredients(),0===e.ingredients.length&&(e.slCurrentPage=e.GetPreviousPage(e.slCurrentPage),e.ShopListServ.Total=e.slcollectionSize,0!==e.slcollectionSize&&e.OnPageChanged(e.slCurrentPage))}),this.WatchIngCle=this.ShopListServ.IngredientClear.subscribe(function(){e.slcollectionSize=0,e.ShopListServ.Total=e.slcollectionSize})}},{key:"GetPreviousPage",value:function(e){return e>1?e-1:1}},{key:"OnPageChanged",value:function(e){var t=this;this.slCurrentPage=e,this.FetchOnInint=this.DataServ.FetchShoppingList(e,o.a.ShoppingListPageSize).subscribe(function(){t.ingredients=t.ShopListServ.GetIngredients(),t.router.navigate(["../",e.toString()],{relativeTo:t.activeroute})})}}]),t}(),k.\u0275fac=function(e){return new(e||k)(c.Ob(a.a),c.Ob(u.a),c.Ob(l.a),c.Ob(u.c))},k.\u0275cmp=c.Ib({type:k,selectors:[["app-shopping-list"]],decls:3,vars:2,consts:[["class","row",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"row"],[1,"col"],[3,"type","close",4,"ngIf"],[1,"list-group","mb-1"],["style","cursor: pointer;","class","list-group-item list-group-item-action d-flex justify-content-between align-items-center",3,"ngClass","click",4,"ngFor","ngForOf"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange",4,"ngIf"],[3,"type","close"],[1,"list-group-item","list-group-item-action","d-flex","justify-content-between","align-items-center",2,"cursor","pointer",3,"ngClass","click"],[1,"badge","badge-success","badge-pill"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(e,t){1&e&&(c.Ac(0,C,8,3,"div",0),c.Tb(1,"div",1),c.Ac(2,y,3,0,"div",2),c.Sb()),2&e&&(c.kc("ngIf",!t.IsLoading),c.Bb(2),c.kc("ngIf",t.IsLoading))},directives:[r.l,f,r.k,v.a,r.j,v.i],styles:[""]}),k),canActivate:[s("dxYa").a],data:{expectedRole:"admin_role_CRUD"}}],w=((T=function t(){e(this,t)}).\u0275mod=c.Mb({type:T}),T.\u0275inj=c.Lb({factory:function(e){return new(e||T)},imports:[[r.c,d.e,v.b,v.j,u.g.forChild(P)]]}),T)}}])}();