<template>
  <transition name="publishAnimate">
    <div :class="'publish publish_' + publishClass" v-if="show">
      <div class="close"><span @click="close">×</span></div>
      <!-- 个人信息 -->
      <div class="userInfo">
        <div class="avatar" :style="{backgroundImage: 'url(' + userStore.avatar + ')'}"></div>
        <div class="infoBox">
          <div class="name">{{ userStore.username }}</div>
          <div class="email">{{ userStore.email }}</div>
        </div>
      </div>
      <!-- 发布区 -->
      <div class="publishBox">
        <textarea autofocus :placeholder="placeholder"></textarea>
      </div>
      <!-- 功能区 -->
      <div class="ribbon">
        <input class="alias" type="text" placeholder="使用化名">
        <div class="btn anonymousBox">是否匿名</div>
        <div class="btn publishBtn">点击发布</div>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { LayoutModeEnum } from "@/constant/enum";

defineProps({
  show: { type: Boolean },
});

const monitorStore = useMonitor();
const processStore = useProcess();
const contextStore = useContext();
const userStore = useUser();
let publishClass = ref(monitorStore.value.layoutMode == LayoutModeEnum.NORMAL ? "normal" : "flip");
let placeholder = ref(`${userStore.value.isLoggedin ? "" : "仅登录用户可发布，"}最多${contextStore.value.messageMaxLen}字。`);

function close() {
  processStore.value.isShowPublish = false;
}

function jumpTo(site: string) {
  window.open(site);
}
</script>

<style scoped lang="scss">
@import "@/assets/scss/global.scss";

$dialogWidthNormal: 500px;
$dialogHeightNormal: 290px;
$dialogWidthFlip: 450px;
$dialogHeightFlip: 300px;
$avatarSize: 50px;
$nameHeight: 30px;

.publishAnimate-enter-active {
  animation: zoomIn 0.5s;
}
.publishAnimate-leave-active {
  animation: zoomOut 0.5s;
}

.publish {
  position: fixed;
  z-index: $zIndexPublish;
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
  .userInfo {
    display: flex;
    justify-content: flex-start;
    .avatar {
      width: $avatarSize;
      height: $avatarSize;
      background-size: 100% 100%;
      border-radius: 50%;
      cursor: pointer;
    }
    .infoBox {
      height: $avatarSize;
      margin-left: 10px;
      .name {
        height: $nameHeight;
        line-height: $nameHeight;
        font-size: 18px;
      }
      .email {
        line-height: $avatarSize - $nameHeight;
        font-size: 14px;
      }
    }
  }
  .publishBox {
    margin-top: 10px;
    textarea {
      width: calc(100% - 20px);
      height: 130px;
      background: none;
      border-radius: 5px;
      border: 2px solid #999;
      resize: none;
      padding: 10px;
      color: $normalColor;
      outline: none;
    }
  }
  .ribbon {
    height: 30px;
    margin-top: 15px;
    display: flex;
    justify-content: flex-end;
    .alias {
      width: 80px;
      height: 30px;
      border-radius: 5px;
      border: 2px solid #999;
      text-align: center;
      color: $normalColor;
      cursor: pointer;
      background: none;
      outline: none;
      font-size: 16px;
      transition: 0.3s;
      &::-webkit-input-placeholder {
        color: $normalColor;
        font-size: 15px;
        transition: 0.3s;
        transform: translateY(1.5px);
      }
      &:-moz-placeholder {
        color: $normalColor;
        font-size: 15px;
        transition: 0.3s;
        transform: translateY(1.5px);
      }
      &::-moz-placeholder {
        color: $normalColor;
        font-size: 15px;
        transition: 0.3s;
        transform: translateY(1.5px);
      } 
      &:-ms-input-placeholder {
        color: $normalColor;
        font-size: 15px;
        transition: 0.3s;
        transform: translateY(1.5px);
      }
      &:hover {
        background: #999;
        color: #fff;
        &::-webkit-input-placeholder { color: #fff; }
        &:-moz-placeholder { color: #fff; }
        &::-moz-placeholder { color: #fff; }
        &:-ms-input-placeholder { color: #fff; }
      }
      &:focus {
        background: none;
        width: 120px;
        color: $normalColor;
        &::-webkit-input-placeholder { color: #999; }
        &:-moz-placeholder { color: #999; }
        &::-moz-placeholder { color: #999; }
        &:-ms-input-placeholder { color: #999; }
      }
    }
    .btn {
      width: 80px;
      height: 30px;
      border-radius: 5px;
      border: 2px solid #999;
      margin-left: 10px;
      text-align: center;
      line-height: 30px;
      cursor: pointer;
      transition: 0.3s;
      &:hover {
        background: #999;
        color: #fff;
      }
    }
  }
}

.publish_normal {
  top: calc(45% - $dialogHeightNormal / 2);
  left: calc(50% - $dialogWidthNormal / 2);
  width: $dialogWidthNormal;
  height: $dialogHeightNormal;
}

.publish_flip {
  top: calc(50% - $dialogHeightFlip / 2);
  left: calc(50% - $dialogWidthFlip / 2);
  transform: rotateZ(90deg);
  width: $dialogWidthFlip;
  height: $dialogHeightFlip;
}
</style>