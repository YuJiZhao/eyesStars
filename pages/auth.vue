<template>
</template>

<script setup lang="ts">
import config from "@/config";

const monitorStore = useMonitor();
const route = useRoute()
const runtimeConfig = useRuntimeConfig();

if (!monitorStore.value.isServer) {
    let sToken = route.query[config.storageKey.sToken] as string;
    let lToken = route.query[config.storageKey.lToken] as string;
    if (sToken != undefined && lToken != undefined) {
        localStorage.setItem(config.storageKey.sToken, sToken);
        localStorage.setItem(config.storageKey.lToken, lToken);
    }
    window.location.replace(runtimeConfig.public.siteUrl);
}
</script>