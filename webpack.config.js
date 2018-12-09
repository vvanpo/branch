/*const clientConfig = {
    entry: './client/index.ts',
}*/

const siteConfig = {
    entry: './site/index.ts',
    output: {
        filename: 'bundle.js',
        path: __dirname + '/dist',
    },
    module: {
        rules: [
            {
                test: /\.ts$/, use: ['ts-loader'],
            },
            {
                test: /\.html$/,
                use: {
                    loader: 'html-loader',
                    options: {
                        attrs: ['script:src'],
                    },
                },
            },
        ],
    },
}

module.exports = (env, argv) => {
    return siteConfig
}
