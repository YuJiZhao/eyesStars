interface MonitorInterface {
    isServer: boolean;
}

export const useMonitor = () => useState<MonitorInterface>("monitor", () => {
    return {
        isServer: true
    }
})