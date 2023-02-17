<template>
  <transition name="starAnimate">
    <div :class="'star star_' + starClass" v-if="show">
      <div class="close"><span @click="close">×</span></div>
      <div class="userInfo">
        <div class="avatar" :style="{backgroundImage: 'url(http://user-cdn.eyescode.top/avatar/stars.jpg)'}"></div>
        <div class="infoBox">
          <div class="name">耶瞳</div>
          <div class="time">2023-02-17 21:52:00</div>
        </div>
      </div>
      <BaseDivider />
      <div class="content">hello world!</div>
      <div class="footer">
        <div class="numbers">序号: 10000</div>
        <div class="views">浏览量: 10000</div>
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
let starClass = ref(monitorStore.value.layoutMode == LayoutModeEnum.NORMAL ? "normal" : "flip");

function close() {
  processStore.value.isShowStar = false;
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

.starAnimate-enter-active {
  animation: zoomIn 0.5s;
}
.starAnimate-leave-active {
  animation: zoomOut 0.5s;
}

.star {
  position: fixed;
  z-index: $zIndexStar;
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
      .time {
        line-height: $avatarSize - $nameHeight;
        font-size: 14px;
      }
    }
  }
  .content {
    margin-top: 10px;
  }
  .footer {
    display: flex;
    position: absolute;
    bottom: 10px;
    right: 10px;
    div {
      margin-left: 10px;
    }
  }
}

.star_normal {
  top: calc(45% - $dialogHeightNormal / 2);
  left: calc(50% - $dialogWidthNormal / 2);
  width: $dialogWidthNormal;
  height: $dialogHeightNormal;
}

.star_flip {
  top: calc(50% - $dialogHeightFlip / 2);
  left: calc(50% - $dialogWidthFlip / 2);
  transform: rotateZ(90deg);
  width: $dialogWidthFlip;
  height: $dialogHeightFlip;
}
</style>