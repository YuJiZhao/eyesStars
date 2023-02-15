<template>
  <NuxtPage v-if="isShow" />
</template>

<script setup lang="ts">
import config from "@/config";
import { useMonitor } from "./composables/useMonitor";
import { initSite } from "@/request/api";
import { LayoutModeEnum } from "@/constant/enum";

const runtimeConfig = useRuntimeConfig();
const monitorStore = useMonitor();
const contextStore = useContext();
let isShow = ref(false);

// 初始化服务信息
monitorStore.value.isServer = "isServer" in runtimeConfig;

// 服务端数据获取，初始化配置信息
if (monitorStore.value.isServer) {
  initSite().then(({ code, msg, data }) => {
    if (code == 200) {
      contextStore.value = data;
    } else {
      console.error("配置初始化错误!赶紧debug!" + msg);
    }
  });
}

/**
 * 判断布局模式
 * @description
 *    判断设备应该使用正常模式还是翻转模式：
 *        正常模式：设备默认布局
 *        翻转模式：翻转屏幕布局
 *    如果屏幕高远大于宽，那么如果还是采用默认布局就太难看了，所以此时得翻转布局，让用户横着看。
 *    因为手机浏览器好像没有提供横屏模式api，所以得自己调整布局实现，有我也懒得搞，因为兼容各个手机厂商太麻烦了
 * @return boolean true：正常模式、false：翻转模式
 */
function determineLayoutMode() {
  // 我更倾向于使用正常模式，因为翻转屏幕好麻烦，容易写出bug
  if (window.innerWidth / window.innerHeight > config.layoutModeThreshold) {
    initNormalMode();
  } else {
    initFlipMode();
  }
}

// 正常模式数据准备
function initNormalMode() {
  monitorStore.value.layoutMode = LayoutModeEnum.NORMAL;
  monitorStore.value.length = window.innerWidth;
  monitorStore.value.height = window.innerHeight;
}

// 翻转模式数据准备
function initFlipMode() {
  monitorStore.value.layoutMode = LayoutModeEnum.Flip;
  monitorStore.value.length = window.innerHeight;
  monitorStore.value.height = window.innerWidth;
}

// 控制台打印应用标签
function printAPPLabel() {
  console.log(
    "\n" +
      ` %c ${contextStore.value.appNameEN} v${contextStore.value.appVersion}` +
      ` %c ${contextStore.value.siteUrl} ` +
      "\n",
    "color: #fadfa3; background: #030307; padding:5px 0;",
    "background: rgb(127,200,248); padding:5px 0;"
  );
}

onMounted(() => {
  determineLayoutMode()
  isShow.value = true;
  printAPPLabel();
});
</script>