
export function template(html: string, style: string): HTMLTemplateElement {
    const tmpl = document.createElement('template')
    tmpl.innerHTML = "<style>" + style + "</style>" + html
    return tmpl
}

export function shadow(el: HTMLElement, tmpl: HTMLTemplateElement) {
    const shadow = el.attachShadow({mode: "open"})
    shadow.appendChild(tmpl.content.cloneNode(true))
}
