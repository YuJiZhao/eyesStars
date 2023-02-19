<template>
  <transition name="noticeAnimate">
    <div :class="'notice notice_' + noticeClass" v-if="show">
      <div class="close"><span @click="close">×</span></div>
      <div class="title">{{ contextStore.appName }}</div>
      <div class="content" v-html="contextStore.notice"></div>
      <div class="btnGroup">
        <div @click="goPublish">发布星星</div>
        <div @click="jumpTo(contextStore.authorSite)">关于作者</div>
        <div @click="jumpTo(contextStore.authorSite)">视频介绍</div>
        <div @click="jumpTo(contextStore.githubAddress)">代码仓库</div>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { LayoutModeEnum } from "@/constant/enum";
import { showDialog } from "@/bussiness/process"; 
import { DialogEnum } from "@/constant/enum";
import config from "~~/config";

defineProps({
  show: { type: Boolean },
});

const monitorStore = useMonitor();
const processStore = useProcess();
const contextStore = useContext();
let dialogDuration = ref(config.dialogDuration);
let noticeClass = ref(
  monitorStore.value.layoutMode == LayoutModeEnum.NORMAL ? "normal" : "flip"
);

function close() {
  processStore.value.isShowNotice = false;
}

function jumpTo(site: string) {
  window.open(site);
}

function goPublish() {
  showDialog(DialogEnum.PUBLISH);
}
</script>

<style scoped lang="scss">
@import "@/assets/scss/global.scss";

$dialogWidthNormal: 500px;
$dialogHeightNormal: 290px;
$dialogWidthFlip: 450px;
$dialogHeightFlip: 300px;

.noticeAnimate-enter-active {
  animation: zoomIn 0.3s;
}
.noticeAnimate-leave-active {
  animation: zoomOut 0.3s;
}

.notice {
  position: fixed;
  z-index: $zIndexNotice;
  background: rgba(#fff, 0.8);
  border-radius: 10px;
  box-shadow: 0 0 5px #fff;
  -webkit-box-shadow: 0 0 5px #fff;
  -moz-box-shadow: 0 0 5px #fff;
  padding: 5px 10px;
  color: $normalColor;
  .close {
    text-align: right;
    span {
      cursor: pointer;
      transition: 0.3s;
      &:hover {
        color: $activeColor;
      }
    }
  }
  .title {
    color: $titleColor;
  }
  .btnGroup {
    display: flex;
    justify-content: space-around;
    margin: 20px 0;
    div {
      cursor: pointer;
      transition: 0.2s;
      &:hover {
        color: $activeColor;
      }
    }
  }
}

.notice_normal {
  top: calc(45% - $dialogHeightNormal / 2);
  left: calc(50% - $dialogWidthNormal / 2);
  width: $dialogWidthNormal;
  height: $dialogHeightNormal;
  .title {
    height: 30px;
    font-size: 20px;
    text-align: center;
    line-height: 30px;
    margin-bottom: 10px;
  }
}

.notice_flip {
  top: calc(50% - $dialogHeightFlip / 2);
  left: calc(50% - $dialogWidthFlip / 2);
  transform: rotateZ(90deg);
  width: $dialogWidthFlip;
  height: $dialogHeightFlip;
  .title {
    height: 30px;
    font-size: 20px;
    text-align: center;
    line-height: 30px;
    margin-bottom: 10px;
  }
}
</style>