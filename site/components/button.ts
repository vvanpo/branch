import html from "./button.html"
import style from "./button.scss"
import { template, shadow } from "./template.ts"

const tmpl = template(html, style)

export default class extends HTMLElement {
    constructor() {
        super()
        shadow(this, tmpl)
    }
}
