import * as zrender from "zrender";
import { decorateStar, moon, slogan } from "@/components/stars/product";
import config from "@/config";

// 绘制静态元素
export let drawStatic = () => {
    let group = new zrender.Group();
    let zlevel = config.staticZLevel;
    
    group.add(decorateStar(zlevel));
    
    return group;
}