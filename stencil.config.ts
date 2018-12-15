import { Config } from "@stencil/core"
import { sass } from "@stencil/sass"

export const config: Config = {
    globalStyle: "site/global/variables.scss",
    namespace: "titian",
    plugins: [sass()],
    srcDir: "site",
}
