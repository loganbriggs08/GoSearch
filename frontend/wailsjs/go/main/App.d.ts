// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {backend} from '../models';

export function CacheSystem():Promise<void>;

export function ClearCache():Promise<boolean>;

export function CloseApp():Promise<void>;

export function GetCurrentPage():Promise<string>;

export function HandleButtonClickEvent(arg1:any):Promise<void>;

export function Search(arg1:string):Promise<Array<backend.FileReturnStruct>>;

export function SetPage(arg1:string):Promise<void>;

export function ToggleFavorite(arg1:string,arg2:string,arg3:boolean):Promise<Array<backend.FileReturnStruct>>;
