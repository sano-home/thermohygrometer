module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: 'tsconfig.json',
    sourceType: 'module'
  },
  plugins: [
    '@typescript-eslint',
    'eslint-plugin-import',
    'eslint-plugin-jsdoc',
    'eslint-plugin-no-null',
    'eslint-plugin-prefer-arrow',
  ],
  settings: {
    react: {
      version: 'detect',
    },
  },
  env: {
    'browser': true,
    'node': true,
    'es6': true,
    'jest': true
  },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/eslint-recommended',
    'plugin:@typescript-eslint/recommended-requiring-type-checking',
    'plugin:@typescript-eslint/recommended',
  ],
  ignorePatterns: [
    'dist/',
    'node_modules/',
  ],
  rules: {
    '@typescript-eslint/no-explicit-any': 'off',
    '@typescript-eslint/explicit-function-return-type': 'off',
    '@typescript-eslint/no-unused-vars': 'off',
    '@typescript-eslint/no-inferrable-types': 'off',
    '@typescript-eslint/await-thenable': 'off',
    '@typescript-eslint/no-use-before-define': 'off',

    '@typescript-eslint/array-type': [
      'error',
      {
        'default': 'array'
      }
    ],
    '@typescript-eslint/ban-types': [
      'error',
      {
        'types': {
          'Object': {
            'message': 'Avoid using the `Object` type. Did you mean `object`?'
          },
          'Function': {
            'message': 'Avoid using the `Function` type. Prefer a specific function type, like `() => void`.'
          },
          'Boolean': {
            'message': 'Avoid using the `Boolean` type. Did you mean `boolean`?'
          },
          'Number': {
            'message': 'Avoid using the `Number` type. Did you mean `number`?'
          },
          'String': {
            'message': 'Avoid using the `String` type. Did you mean `string`?'
          },
          'Symbol': {
            'message': 'Avoid using the `Symbol` type. Did you mean `symbol`?'
          }
        }
      }
    ],
    '@typescript-eslint/consistent-type-definitions': 'off',
    '@typescript-eslint/dot-notation': 'error',
    '@typescript-eslint/explicit-member-accessibility': [
      'off',
      {
        'accessibility': 'explicit'
      }
    ],
    '@typescript-eslint/indent': 'off',
    '@typescript-eslint/interface-name-prefix': 'off',
    '@typescript-eslint/member-delimiter-style': [
      'error',
      {
        'multiline': {
          'delimiter': 'semi',
          'requireLast': true
        },
        'singleline': {
          'delimiter': 'semi',
          'requireLast': false
        }
      }
    ],
    '@typescript-eslint/member-ordering': 'error',
    '@typescript-eslint/no-empty-interface': 'off',
    'no-param-reassign': 'error',
    '@typescript-eslint/no-parameter-properties': 'off',
    '@typescript-eslint/no-require-imports': 'off',
    '@typescript-eslint/no-unnecessary-boolean-literal-compare': 'off',
    '@typescript-eslint/no-unused-expressions': 'error',
    '@typescript-eslint/prefer-for-of': 'error',
    '@typescript-eslint/prefer-function-type': 'error',
    '@typescript-eslint/quotes': [
      'error',
      'single'
    ],
    '@typescript-eslint/semi': [
      'error',
      'always'
    ],
    '@typescript-eslint/strict-boolean-expressions': 'off',
    '@typescript-eslint/triple-slash-reference': [
      'error',
      {
        'path': 'always',
        'types': 'prefer-import',
        'lib': 'always'
      }
    ],
    '@typescript-eslint/type-annotation-spacing': 'off',
    '@typescript-eslint/unified-signatures': 'error',
    'arrow-body-style': 'off',
    'arrow-parens': [
      'off',
      'always'
    ],
    'brace-style': [
      'off',
      'off'
    ],
    'camelcase': 'error',
    'complexity': 'off',
    'constructor-super': 'error',
    'curly': [
      'error',
      'multi-line'
    ],
    'eol-last': 'off',
    'eqeqeq': [
      'error',
      'smart'
    ],
    'guard-for-in': 'error',
    'id-match': 'error',
    'import/no-unassigned-import': ['error', {
      allow: [
        '@testing-library/jest-dom/extend-expect',
        'mutationobserver-shim',
        '**/*.css',
      ]
    }],
    'import/order': [
      'error',
      {
        'newlines-between': 'always-and-inside-groups',
        groups: ['builtin', 'external', 'unknown', 'internal', ['parent', 'sibling', 'index']],
        pathGroups: [
          {
            group: 'internal',
            pattern: 'containers/**',
          },
          {
            group: 'internal',
            pattern: 'components/**',
          },
          {
            group: 'internal',
            pattern: 'routeComponents/**',
          },
          {
            group: 'parent',
            pattern: '../**',
          },
          {
            group: 'sibling',
            pattern: './**',
          },
        ],
        pathGroupsExcludedImportTypes: ['builtin', 'external', 'index'],
      }
    ],
    'jsdoc/check-alignment': 'error',
    'jsdoc/check-indentation': 'error',
    'jsdoc/newline-after-description': 'error',
    'linebreak-style': 'off',
    'max-classes-per-file': [
      'error',
      1
    ],
    'max-len': 'off',
    'new-parens': 'off',
    'newline-per-chained-call': 'off',
    'no-bitwise': 'error',
    'no-caller': 'error',
    'no-cond-assign': 'error',
    'no-console': 'error',
    'no-debugger': 'error',
    'no-duplicate-imports': 'error',
    'no-empty': 'error',
    'no-eval': 'error',
    'no-extra-semi': 'off',
    'no-fallthrough': 'off',
    'no-invalid-this': 'error',
    'no-irregular-whitespace': 'off',
    'no-multiple-empty-lines': 'off',
    'no-new-wrappers': 'error',
    'no-null/no-null': 'off',
    'no-restricted-imports': [
      'error',
      'lodash'
    ],
    'no-shadow': [
      'error',
      {
        'hoist': 'all'
      }
    ],
    'no-throw-literal': 'error',
    'no-trailing-spaces': 'error',
    'no-undef-init': 'error',
    'no-underscore-dangle': 'error',
    'no-unsafe-finally': 'error',
    'no-unused-labels': 'error',
    'object-shorthand': 'off',
    'one-var': [
      'off',
      'never'
    ],
    'padding-line-between-statements': [
      'off',
      {
        'blankLine': 'always',
        'prev': '*',
        'next': 'return'
      }
    ],
    'prefer-arrow/prefer-arrow-functions': ['error', {
      allowStandaloneDeclarations: true
    }],
    'prefer-template': 'error',
    'quote-props': 'off',
    'radix': 'error',
    'space-before-function-paren': 'off',
    'space-in-parens': [
      'off',
      'never'
    ],
    'spaced-comment': [
      'error',
      'always',
      {
        'markers': [
          '/'
        ]
      }
    ],
    'use-isnan': 'error',
    'valid-typeof': 'off',
    'no-plusplus': 'error',
    'no-else-return': 'error',
    'no-array-constructor': 'error',
    '@typescript-eslint/typedef': ['error', {
      parameter: true,
      propertyDeclaration: true,
      memberVariableDeclaration: false,
      arrowParameter: false,
    }],
    'no-new-func': 'error',
    'no-else-return': 'error',
    'prefer-arrow-callback': 'error',
  }
};
