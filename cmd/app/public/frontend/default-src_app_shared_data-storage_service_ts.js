"use strict";
(self["webpackChunkpainting_shop_front"] = self["webpackChunkpainting_shop_front"] || []).push([["default-src_app_shared_data-storage_service_ts"],{

/***/ 162:
/*!********************************************!*\
  !*** ./src/app/admin/media/media.model.ts ***!
  \********************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "FiLe": () => (/* binding */ FiLe),
/* harmony export */   "FilesResponse": () => (/* binding */ FilesResponse)
/* harmony export */ });
class FiLe {
    constructor(fname, fsize, ftype, flink, plink, ID) {
        this.Filename = fname;
        this.Filesize = fsize;
        this.Filetype = ftype;
        this.FileID = flink;
        this.ID = ID;
        this.PreviewID = plink;
    }
}
class FilesResponse {
}


/***/ }),

/***/ 2245:
/*!**********************************************!*\
  !*** ./src/app/admin/media/media.service.ts ***!
  \**********************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "MediaService": () => (/* binding */ MediaService)
/* harmony export */ });
/* harmony import */ var _media_model__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./media.model */ 162);
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! rxjs */ 228);
/* harmony import */ var src_environments_environment__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! src/environments/environment */ 2340);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/core */ 3184);




class MediaService {
    constructor() {
        this.FileSelected = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.FilesUpdated = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.FilesInserted = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.FilesDeleted = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.FilesChanged = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.Files = [];
    }
    GetFiles() {
        return this.Files.slice();
    }
    SetFiles(newfiles) {
        this.Files = newfiles;
        this.FilesUpdated.next();
    }
    SetPagination(Total, Limit, Offset) {
        this.Total = Total;
    }
    SelectItemFilesList(f) {
        this.CurrentSelectedItem = f;
        this.FileSelected.next(f);
    }
    IsCurrentSelected(user) {
        return this.CurrentSelectedItem === user;
    }
    GetFileById(id) {
        if (id < this.Files.length && id > 0) {
            return this.Files[id];
        }
        else {
            return this.Files[0];
        }
    }
    UpdateExistingFile(f, Index) {
        this.Files[Index] = f;
        this.FilesChanged.next(f);
    }
    AddNewFile(f) {
        const nf = new _media_model__WEBPACK_IMPORTED_MODULE_0__.FiLe(f.Filename, f.Filesize, f.Filetype, f.FileID, f.PreviewID, f.ID);
        if (this.Files.length < src_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.AdminUserListPageSize) {
            this.Files.push(nf);
        }
        this.FilesChanged.next(nf);
        this.FilesInserted.next();
    }
    DeleteFile(Index) {
        this.Files.splice(Index, 1);
        this.FilesDeleted.next();
    }
}
MediaService.ɵfac = function MediaService_Factory(t) { return new (t || MediaService)(); };
MediaService.ɵprov = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdefineInjectable"]({ token: MediaService, factory: MediaService.ɵfac, providedIn: 'root' });


/***/ }),

/***/ 9101:
/*!****************************************************!*\
  !*** ./src/app/admin/sessions/sessions.service.ts ***!
  \****************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "SessionsService": () => (/* binding */ SessionsService)
/* harmony export */ });
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! rxjs */ 228);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ 3184);


class SessionsService {
    constructor() {
        this.SessionsUpdated = new rxjs__WEBPACK_IMPORTED_MODULE_0__.Subject();
        this.SessionsInserted = new rxjs__WEBPACK_IMPORTED_MODULE_0__.Subject();
        this.SessionsDeleted = new rxjs__WEBPACK_IMPORTED_MODULE_0__.Subject();
        this.SessionsChanged = new rxjs__WEBPACK_IMPORTED_MODULE_0__.Subject();
        this.SessionsSelected = new rxjs__WEBPACK_IMPORTED_MODULE_0__.Subject();
        this.Sessions = [];
    }
    GetSessions() {
        return this.Sessions.slice();
    }
    SetSessions(newsessions) {
        this.Sessions = newsessions;
        this.SessionsUpdated.next();
    }
    SetPagination(Total, Limit, Offset) {
        this.Total = Total;
    }
    SelectItemSessionsList(s) {
        this.CurrentSelectedItem = s;
        this.SessionsSelected.next(s);
    }
    IsCurrentSelected(s) {
        return this.CurrentSelectedItem === s;
    }
    GetSessionById(id) {
        if (id < this.Sessions.length && id > 0) {
            return this.Sessions[id];
        }
        else {
            return this.Sessions[0];
        }
    }
    DeleteSession(Index) {
        this.Sessions.splice(Index, 1);
        this.SessionsDeleted.next();
    }
}
SessionsService.ɵfac = function SessionsService_Factory(t) { return new (t || SessionsService)(); };
SessionsService.ɵprov = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdefineInjectable"]({ token: SessionsService, factory: SessionsService.ɵfac, providedIn: 'root' });


/***/ }),

/***/ 6729:
/*!********************************************!*\
  !*** ./src/app/admin/users/users.model.ts ***!
  \********************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "User": () => (/* binding */ User),
/* harmony export */   "UsersResponse": () => (/* binding */ UsersResponse)
/* harmony export */ });
class User {
    constructor(Role, Email, Phone, Name) {
        this.GUID = '';
        this.Role = Role;
        this.Email = Email;
        this.Phone = Phone;
        this.Name = Name;
        this.IsAdmin = false;
        this.Confirmed = false;
        this.SecondFactor = false;
        this.Disabled = false;
    }
}
class UsersResponse {
}


/***/ }),

/***/ 9112:
/*!**********************************************!*\
  !*** ./src/app/admin/users/users.service.ts ***!
  \**********************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "UsersService": () => (/* binding */ UsersService)
/* harmony export */ });
/* harmony import */ var _users_model__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./users.model */ 6729);
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! rxjs */ 228);
/* harmony import */ var src_environments_environment__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! src/environments/environment */ 2340);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/core */ 3184);




class UsersService {
    constructor() {
        this.UserSelected = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.UsersUpdated = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.UsersInserted = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.UsersDeleted = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.UsersChanged = new rxjs__WEBPACK_IMPORTED_MODULE_2__.Subject();
        this.Users = [];
    }
    GetUsers() {
        return this.Users.slice();
    }
    SetUsers(newusers) {
        this.Users = newusers;
        this.UsersUpdated.next();
    }
    SetPagination(Total, Limit, Offset) {
        this.Total = Total;
    }
    SelectItemUsersList(ingredient) {
        this.CurrentSelectedItem = ingredient;
        this.UserSelected.next(ingredient);
    }
    IsCurrentSelected(user) {
        return this.CurrentSelectedItem === user;
    }
    GetUserById(id) {
        if (id < this.Users.length && id > 0) {
            return this.Users[id];
        }
        else {
            return this.Users[0];
        }
    }
    UpdateExistingUser(UserToUpdate, Index) {
        this.Users[Index] = UserToUpdate;
        this.UsersChanged.next(UserToUpdate);
    }
    AddNewUser(NewUser) {
        const NewUserToAdd = new _users_model__WEBPACK_IMPORTED_MODULE_0__.User(NewUser.Role, NewUser.Email, NewUser.Phone, NewUser.Name);
        if (this.Users.length < src_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.AdminUserListPageSize) {
            this.Users.push(NewUserToAdd);
        }
        this.UsersChanged.next(NewUserToAdd);
        this.UsersInserted.next();
    }
    DeleteUser(Index) {
        this.Users.splice(Index, 1);
        this.UsersDeleted.next();
    }
}
UsersService.ɵfac = function UsersService_Factory(t) { return new (t || UsersService)(); };
UsersService.ɵprov = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdefineInjectable"]({ token: UsersService, factory: UsersService.ɵfac, providedIn: 'root' });


/***/ }),

/***/ 3649:
/*!************************************************!*\
  !*** ./src/app/shared/data-storage.service.ts ***!
  \************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "DataStorageService": () => (/* binding */ DataStorageService)
/* harmony export */ });
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @angular/common/http */ 8784);
/* harmony import */ var _shared_shared_model__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../shared/shared.model */ 3481);
/* harmony import */ var rxjs_operators__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! rxjs/operators */ 9337);
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../../environments/environment */ 2340);
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! rxjs */ 228);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _admin_users_users_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../admin/users/users.service */ 9112);
/* harmony import */ var _admin_media_media_service__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../admin/media/media.service */ 2245);
/* harmony import */ var _admin_sessions_sessions_service__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../admin/sessions/sessions.service */ 9101);










class DataStorageService {
    constructor(http, users, media, sessions) {
        this.http = http;
        this.users = users;
        this.media = media;
        this.sessions = sessions;
        this.LoadingData = new rxjs__WEBPACK_IMPORTED_MODULE_5__.Subject();
        this.RecivedError = new rxjs__WEBPACK_IMPORTED_MODULE_5__.Subject();
        this.PaginationSet = new rxjs__WEBPACK_IMPORTED_MODULE_5__.Subject();
        this.FileUploadProgress = new rxjs__WEBPACK_IMPORTED_MODULE_5__.Subject();
        this.FileUploaded = new rxjs__WEBPACK_IMPORTED_MODULE_5__.Subject();
        this.UserUpdateInsert = new rxjs__WEBPACK_IMPORTED_MODULE_5__.Subject();
        this.CurrentUserFetch = new rxjs__WEBPACK_IMPORTED_MODULE_5__.Subject();
        this.TwoFactorSub = new rxjs__WEBPACK_IMPORTED_MODULE_5__.Subject();
    }
    FetchFilesList(page, limit) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Page: page.toString(),
                Limit: limit.toString()
            })
        };
        return this.http
            .get(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetFileUrl, httpOptions)
            .pipe((0,rxjs_operators__WEBPACK_IMPORTED_MODULE_7__.tap)(recresp => {
            this.media.SetFiles(recresp.Files);
            this.media.SetPagination(recresp.Total, recresp.Limit, recresp.Offset);
            this.LoadingData.next(false);
        }, (error) => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        }));
    }
    FileUpload(FileToUpload) {
        const formdatafile = new FormData();
        formdatafile.append('file', FileToUpload, FileToUpload.name);
        this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetFileUrl, formdatafile, {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({}),
            reportProgress: true,
            observe: 'events'
        }).subscribe((curevent) => {
            if (curevent.type === _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpEventType.UploadProgress) {
                this.FileUploadProgress.next(String(curevent.loaded / curevent.total * 100));
            }
            else if (curevent.type === _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpEventType.Response) {
                if (curevent.ok) {
                    this.FileUploaded.next(curevent.body);
                }
            }
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    DeleteFile(FileID, NoMessage) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                FileID: FileID.toString()
            })
        };
        this.http.delete(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetFileUrl, httpOptions)
            .subscribe(response => {
            if (!NoMessage) {
                this.RecivedError.next(response);
            }
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    FetchSessionsList(page, limit) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Page: page.toString(),
                Limit: limit.toString()
            })
        };
        return this.http
            .get(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetSessionsUrl, httpOptions)
            .pipe((0,rxjs_operators__WEBPACK_IMPORTED_MODULE_7__.tap)(recresp => {
            this.sessions.SetSessions(recresp.Sessions);
            this.sessions.SetPagination(recresp.Total, recresp.Limit, recresp.Offset);
            this.LoadingData.next(false);
        }, (error) => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        }));
    }
    DeleteSessionByToken(token) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Token: token
            })
        };
        this.http.delete(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetSessionsUrl, httpOptions)
            .subscribe(response => {
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    DeleteSessionByEmail(email) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Email: email
            })
        };
        this.http.delete(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetSessionsUrl, httpOptions)
            .subscribe(response => {
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    FetchUsersList(page, limit) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Page: page.toString(),
                Limit: limit.toString()
            })
        };
        return this.http
            .get(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetUsersUrl, httpOptions)
            .pipe((0,rxjs_operators__WEBPACK_IMPORTED_MODULE_7__.tap)(recresp => {
            this.users.SetUsers(recresp.Users);
            this.users.SetPagination(recresp.Total, recresp.Limit, recresp.Offset);
            this.LoadingData.next(false);
        }, (error) => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        }));
    }
    FetchCurrentUser() {
        this.LoadingData.next(true);
        return this.http.get(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetCurrentUserUrl)
            .subscribe(response => {
            this.CurrentUserFetch.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    SaveCurrentUser(ItemToSave, ChangePassword, NewPassword) {
        this.LoadingData.next(true);
        if (ItemToSave.GUID.length === 0) {
            ItemToSave.GUID = '00000000-0000-0000-0000-000000000000';
        }
        this.GetObsForSaveCurrentUser(ItemToSave, ChangePassword, NewPassword)
            .subscribe(response => {
            this.UserUpdateInsert.next(response);
            this.RecivedError.next(new _shared_shared_model__WEBPACK_IMPORTED_MODULE_0__.ErrorResponse(200, 'Данные сохранены'));
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    SaveUser(ItemToSave, ChangePassword, NewPassword) {
        this.LoadingData.next(true);
        if (ItemToSave.GUID.length === 0) {
            ItemToSave.GUID = '00000000-0000-0000-0000-000000000000';
        }
        this.GetObsForSaveUser(ItemToSave, ChangePassword, NewPassword)
            .subscribe(response => {
            this.UserUpdateInsert.next(response);
            this.RecivedError.next(new _shared_shared_model__WEBPACK_IMPORTED_MODULE_0__.ErrorResponse(200, 'Данные сохранены'));
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    LinkTwoFactor(Key, CurUser) {
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Passcode: Key
            })
        };
        this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.TOTPSettingsUrl, CurUser, httpOptions)
            .subscribe(response => {
            CurUser.SecondFactor = true;
            this.TwoFactorSub.next(CurUser);
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    UnlinkTwoFactor(CurUser) {
        this.http.delete(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.TOTPSettingsUrl)
            .subscribe(response => {
            CurUser.SecondFactor = false;
            this.TwoFactorSub.next(CurUser);
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    GetObsForSaveCurrentUser(ItemToSave, ChangePassword, NewPassword) {
        if (ChangePassword) {
            const httpOptions = {
                headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                    NewPassword: encodeURI(NewPassword)
                })
            };
            return this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetCurrentUserUrl, ItemToSave, httpOptions);
        }
        else {
            return this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetCurrentUserUrl, ItemToSave);
        }
    }
    GetObsForSaveUser(ItemToSave, ChangePassword, NewPassword) {
        if (ChangePassword) {
            const httpOptions = {
                headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                    NewPassword: encodeURI(NewPassword)
                })
            };
            return this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetUsersUrl, ItemToSave, httpOptions);
        }
        else {
            return this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetUsersUrl, ItemToSave);
        }
    }
    DeleteUser(UserToDelete) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                UserID: encodeURI(UserToDelete.GUID)
            })
        };
        this.http.delete(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.GetSetUsersUrl, httpOptions)
            .subscribe(response => {
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    ConfirmEmail(UniqueToken) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Token: UniqueToken
            })
        };
        this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.ConfirmEmailUrl, null, httpOptions)
            .subscribe(response => {
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    SendEmailConfirmEmail(EmailToSend) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Email: EmailToSend
            })
        };
        this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.ResendEmailUrl, null, httpOptions)
            .subscribe(response => {
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    SendEmailResetPassword(EmailToSend) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Email: EmailToSend
            })
        };
        this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.SendEmailResetPassUrl, null, httpOptions)
            .subscribe(response => {
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
    SubmitNewPassword(UniqueToken, NewPass) {
        this.LoadingData.next(true);
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpHeaders({
                Token: UniqueToken,
                NewPassword: encodeURI(NewPass)
            })
        };
        this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.ResetPasswordUrl, null, httpOptions)
            .subscribe(response => {
            this.RecivedError.next(response);
            this.LoadingData.next(false);
        }, error => {
            const errresp = error.error;
            this.RecivedError.next(errresp);
            this.LoadingData.next(false);
        });
    }
}
DataStorageService.ɵfac = function DataStorageService_Factory(t) { return new (t || DataStorageService)(_angular_core__WEBPACK_IMPORTED_MODULE_8__["ɵɵinject"](_angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpClient), _angular_core__WEBPACK_IMPORTED_MODULE_8__["ɵɵinject"](_admin_users_users_service__WEBPACK_IMPORTED_MODULE_2__.UsersService), _angular_core__WEBPACK_IMPORTED_MODULE_8__["ɵɵinject"](_admin_media_media_service__WEBPACK_IMPORTED_MODULE_3__.MediaService), _angular_core__WEBPACK_IMPORTED_MODULE_8__["ɵɵinject"](_admin_sessions_sessions_service__WEBPACK_IMPORTED_MODULE_4__.SessionsService)); };
DataStorageService.ɵprov = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_8__["ɵɵdefineInjectable"]({ token: DataStorageService, factory: DataStorageService.ɵfac, providedIn: 'root' });


/***/ }),

/***/ 3481:
/*!****************************************!*\
  !*** ./src/app/shared/shared.model.ts ***!
  \****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "BackendError": () => (/* binding */ BackendError),
/* harmony export */   "ErrorResponse": () => (/* binding */ ErrorResponse),
/* harmony export */   "Ingredient": () => (/* binding */ Ingredient),
/* harmony export */   "Pagination": () => (/* binding */ Pagination),
/* harmony export */   "ShoppingListResponse": () => (/* binding */ ShoppingListResponse)
/* harmony export */ });
class Ingredient {
    constructor(Name, Amount) {
        this.Name = Name;
        this.Amount = Amount;
    }
}
class ShoppingListResponse {
}
class ErrorResponse {
    constructor(Code, Message) {
        this.Error = new BackendError(Code, Message);
    }
}
class BackendError {
    constructor(Code, Message) {
        this.Code = Code;
        this.Message = Message;
    }
}
class Pagination {
    constructor(Total, Limit, Offset) {
        this.Total = Total;
        this.Limit = Limit;
        this.Offset = Offset;
    }
}


/***/ })

}]);
//# sourceMappingURL=default-src_app_shared_data-storage_service_ts.js.map