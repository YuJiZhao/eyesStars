import { ProcessInterface } from "@/composables/useProcess";
import { DialogEnum } from "@/constant/enum";
import config from "@/config";

// 控制器初始化
let process: null | ProcessInterface = null;
export function initProcess(processStore: ProcessInterface) {
    process = processStore;
}

// 打开弹窗
export function showDialog(dialog: DialogEnum) {
    process!.isShowNotice = false;
    process!.isShowPublish = false;
    process!.isShowStar = false;
    switch (dialog) {
        case DialogEnum.NOTICE:
            process!.isShowNotice = true;
            break;
        case DialogEnum.PUBLISH:
            process!.isShowPublish = true;
            break;
        case DialogEnum.STAR:
            process!.isShowStar = true;
            break;
    }
}

// 打开tip
export function showTip(msg: string) {
    process!.tipMsg = msg;
    process!.isShowTip = true;
}

// 切换星星弹窗
export function switchStar() {
    process!.isShowNotice = false;
    process!.isShowPublish = false;
    if (process!.isShowStar) {
        process!.isShowStar = false;
        setTimeout(() => {
            process!.isShowStar = true;
        }, config.dialogDuration);
    } else {
        process!.isShowStar = true;
    }
}