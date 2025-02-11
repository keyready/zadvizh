import globals from 'globals';
import pluginJs from '@eslint/js';
import importPlugin from 'eslint-plugin-import';
import pluginReact from 'eslint-plugin-react';
import tseslint from 'typescript-eslint';

/** @type {import('eslint').Linter.Config[]} */
export default [
    { files: ['**/*.{js,ts,jsx,tsx}'] },
    { languageOptions: { globals: globals.browser } },
    pluginJs.configs.recommended,
    importPlugin.flatConfigs.recommended,
    ...tseslint.configs.recommended,
    pluginReact.configs.flat.recommended,
    {
        rules: {
            'react/react-in-jsx-scope': 'off',
            'max-len': ['error', { ignoreComments: true, code: 150 }],
            'react/jsx-max-props-per-line': [`error`, { maximum: 5 }],
            'import/no-unresolved': 'off',
            '@typescript-eslint/no-explicit-any': 'off',
            'import/order': [
                'error',
                {
                    'groups': ['builtin', 'external', 'parent', 'sibling', 'index'],
                    'newlines-between': 'always',
                },
            ],
        },
    },
];
