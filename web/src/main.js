import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import createHljsTheme from '@kangc/v-md-editor/lib/theme/hljs';
import hljs from 'highlight.js';

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

const hljsTheme = createHljsTheme({
    Hljs: hljs,
});
hljsTheme.extend((md) => {
    // md为 markdown-it 实例，可以在此处进行修改配置,并使用 plugin 进行语法扩展
    // md.set(option).use(plugin);
});
VueMarkdownEditor.vMdParser.theme(hljsTheme);
const app = createApp(App)

app.use(router).use(VueMarkdownEditor).use(ElementPlus)

app.mount('#app')

