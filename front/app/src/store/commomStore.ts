import {makeAutoObservable} from "mobx";

export class CommonStore {
    lang: string = 'en'
    globalLoading: boolean = false
    footerShow: boolean = true


    constructor() {
        makeAutoObservable(this)
    }


    setGlobalLoading(payload: boolean) {
        this.globalLoading = payload
    }


}