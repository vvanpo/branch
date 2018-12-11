import { Component, Prop } from "@stencil/core"

@Component({
    tag: "titian-button",
    styleUrl: "button.scss",
})
export class Button {

    @Prop() public disabled: boolean
    @Prop() public href: string
    @Prop() public name: string
    @Prop() public reset: boolean
    @Prop() public value: string

    public render() {
        const Tag: string = this.href ? "a" : "button"
        return (
            <Tag {...this.attrs()}>
                <slot />
            </Tag>
        )
    }

    // Determine button or anchor element attributes
    private attrs() {
        const a: any = {}

        if (this.href) {
            a.href = this.href
        } else {
            a.type = this.reset ? "reset" : "button"

            if (this.value) {
                a.name = this.name
                a.type = "submit"
                a.value = this.value
            }

            if (this.disabled) {
                a.disabled = true
            }
        }

        return a
    }
}
