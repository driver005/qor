(function (factory) {
    if (typeof define === 'function' && define.amd) {
        // AMD. Register as anonymous module.
        define(['jquery'], factory);
    } else if (typeof exports === 'object') {
        // Node / CommonJS
        factory(require('jquery'));
    } else {
        // Browser globals.
        factory(jQuery);
    }
})(function ($) {

    'use strict';
    var $document = $(document);
    var NAMESPACE = 'qor.slug';
    var EVENT_ENABLE = 'enable.' + NAMESPACE;
    var EVENT_DISABLE = 'disable.' + NAMESPACE;
    var EVENT_CLICK = 'click.' + NAMESPACE;
    var EVENT_KEYUP = 'keyup.' + NAMESPACE;

    var CLASS_INPUT = '.qor-slug-input';
    var CLASS_SHOW = '.qor-slug-show';
    var CLASS_TRIGGER = '.qor-slug-trigger';

    function QorSlug(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorSlug.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorSlug.prototype = {
        constructor: QorSlug,

        init: function () {
            var $element = this.$element;

            this.$show = $element.find(CLASS_SHOW);
            this.$input = $element.find(CLASS_INPUT);
            this.$trigger = $element.find(CLASS_TRIGGER);
            this.originalShowValue = this.$show.val();

            this.charMap = QorSlug.CHARMAP;

            this.needSyncSlug = this.$trigger.is(':checked');
            this.regChar = '';

            window._.map(this.charMap, function (num, key) {
                this.regChar = this.regChar + key;
            }.bind(this));

            this.regChar = '[' + this.regChar + ']';
            this.regChar = new RegExp(this.regChar, 'g');

            this.bind();
        },

        bind: function () {
            $document
                .on(EVENT_KEYUP, CLASS_INPUT, $.proxy(this.change, this))
                .on(EVENT_CLICK, CLASS_TRIGGER, $.proxy(this.click, this));
        },

        unbind: function () {
            $document
                .off(EVENT_KEYUP, CLASS_INPUT, $.proxy(this.change, this))
                .off(EVENT_CLICK, CLASS_TRIGGER, $.proxy(this.click, this));
        },

        click: function () {
            this.needSyncSlug = this.$trigger.is(':checked');

            if (this.needSyncSlug) {
                this.change();
            } else {
                this.$show.val(this.originalShowValue);
            }

        },

        change: function () {
            if (!this.needSyncSlug) {
                return;
            }

            var inputValue;

            var regSpace = /[-\s]+/g;
            var regChinese = /[\u4e00-\u9fa5]/g;
            var regSymbol = /[#,:;@+!?\\\/]/g;
            var regChar;

            inputValue = this.$input.val();

            var char = inputValue.match(this.regChar);

            if (char && char.length) {
                window._.each(char, function (item) {
                    regChar = new RegExp('[' + item + ']', 'g');
                    inputValue = inputValue.replace(regChar, this.charMap[item]);
                }.bind(this));
            }

            inputValue = inputValue.replace(regChinese, '').replace(regSymbol, '').replace(regSpace, '-');

            this.$show.val(inputValue.toLowerCase());
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }

    };

    QorSlug.DEFAULTS = {};

    QorSlug.CHARMAP = {
        // latin
        '??': 'A',
        '??': 'A',
        '??': 'A',
        '??': 'A',
        '??': 'A',
        '??': 'A',
        '??': 'AE',
        '??': 'C',
        '??': 'E',
        '??': 'E',
        '??': 'E',
        '??': 'E',
        '??': 'I',
        '??': 'I',
        '??': 'I',
        '??': 'I',
        '??': 'D',
        '??': 'N',
        '??': 'O',
        '??': 'O',
        '??': 'O',
        '??': 'O',
        '??': 'O',
        '??': 'O',
        '??': 'O',
        '??': 'U',
        '??': 'U',
        '??': 'U',
        '??': 'U',
        '??': 'U',
        '??': 'Y',
        '??': 'TH',
        '??': 'ss',
        '??': 'a',
        '??': 'a',
        '??': 'a',
        '??': 'a',
        '??': 'a',
        '??': 'a',
        '??': 'ae',
        '??': 'c',
        '??': 'e',
        '??': 'e',
        '??': 'e',
        '??': 'e',
        '??': 'i',
        '??': 'i',
        '??': 'i',
        '??': 'i',
        '??': 'd',
        '??': 'n',
        '??': 'o',
        '??': 'o',
        '??': 'o',
        '??': 'o',
        '??': 'o',
        '??': 'o',
        '??': 'o',
        '??': 'u',
        '??': 'u',
        '??': 'u',
        '??': 'u',
        '??': 'u',
        '??': 'y',
        '??': 'th',
        '??': 'y',
        '???': 'SS',
        // greek
        '??': 'a',
        '??': 'b',
        '??': 'g',
        '??': 'd',
        '??': 'e',
        '??': 'z',
        '??': 'h',
        '??': '8',
        '??': 'i',
        '??': 'k',
        '??': 'l',
        '??': 'm',
        '??': 'n',
        '??': '3',
        '??': 'o',
        '??': 'p',
        '??': 'r',
        '??': 's',
        '??': 't',
        '??': 'y',
        '??': 'f',
        '??': 'x',
        '??': 'ps',
        '??': 'w',
        '??': 'a',
        '??': 'e',
        '??': 'i',
        '??': 'o',
        '??': 'y',
        '??': 'h',
        '??': 'w',
        '??': 's',
        '??': 'i',
        '??': 'y',
        '??': 'y',
        '??': 'i',
        '??': 'A',
        '??': 'B',
        '??': 'G',
        '??': 'D',
        '??': 'E',
        '??': 'Z',
        '??': 'H',
        '??': '8',
        '??': 'I',
        '??': 'K',
        '??': 'L',
        '??': 'M',
        '??': 'N',
        '??': '3',
        '??': 'O',
        '??': 'P',
        '??': 'R',
        '??': 'S',
        '??': 'T',
        '??': 'Y',
        '??': 'F',
        '??': 'X',
        '??': 'PS',
        '??': 'W',
        '??': 'A',
        '??': 'E',
        '??': 'I',
        '??': 'O',
        '??': 'Y',
        '??': 'H',
        '??': 'W',
        '??': 'I',
        '??': 'Y',
        // turkish
        '??': 's',
        '??': 'S',
        '??': 'i',
        '??': 'I',
        '??': 'g',
        '??': 'G',
        // russian
        '??': 'a',
        '??': 'b',
        '??': 'v',
        '??': 'g',
        '??': 'd',
        '??': 'e',
        '??': 'yo',
        '??': 'zh',
        '??': 'z',
        '??': 'i',
        '??': 'j',
        '??': 'k',
        '??': 'l',
        '??': 'm',
        '??': 'n',
        '??': 'o',
        '??': 'p',
        '??': 'r',
        '??': 's',
        '??': 't',
        '??': 'u',
        '??': 'f',
        '??': 'h',
        '??': 'c',
        '??': 'ch',
        '??': 'sh',
        '??': 'sh',
        '??': 'u',
        '??': 'y',
        '??': '',
        '??': 'e',
        '??': 'yu',
        '??': 'ya',
        '??': 'A',
        '??': 'B',
        '??': 'V',
        '??': 'G',
        '??': 'D',
        '??': 'E',
        '??': 'Yo',
        '??': 'Zh',
        '??': 'Z',
        '??': 'I',
        '??': 'J',
        '??': 'K',
        '??': 'L',
        '??': 'M',
        '??': 'N',
        '??': 'O',
        '??': 'P',
        '??': 'R',
        '??': 'S',
        '??': 'T',
        '??': 'U',
        '??': 'F',
        '??': 'H',
        '??': 'C',
        '??': 'Ch',
        '??': 'Sh',
        '??': 'Sh',
        '??': 'U',
        '??': 'Y',
        '??': '',
        '??': 'E',
        '??': 'Yu',
        '??': 'Ya',
        // ukranian
        '??': 'Ye',
        '??': 'I',
        '??': 'Yi',
        '??': 'G',
        '??': 'ye',
        '??': 'i',
        '??': 'yi',
        '??': 'g',
        // czech
        '??': 'c',
        '??': 'd',
        '??': 'e',
        '??': 'n',
        '??': 'r',
        '??': 's',
        '??': 't',
        '??': 'u',
        '??': 'z',
        '??': 'C',
        '??': 'D',
        '??': 'E',
        '??': 'N',
        '??': 'R',
        '??': 'S',
        '??': 'T',
        '??': 'U',
        '??': 'Z',
        // polish
        '??': 'a',
        '??': 'c',
        '??': 'e',
        '??': 'l',
        '??': 'n',
        '??': 's',
        '??': 'z',
        '??': 'z',
        '??': 'A',
        '??': 'C',
        '??': 'E',
        '??': 'L',
        '??': 'N',
        '??': 'S',
        '??': 'Z',
        '??': 'Z',
        // latvian
        '??': 'a',
        '??': 'e',
        '??': 'g',
        '??': 'i',
        '??': 'k',
        '??': 'l',
        '??': 'n',
        '??': 'u',
        '??': 'A',
        '??': 'E',
        '??': 'G',
        '??': 'I',
        '??': 'K',
        '??': 'L',
        '??': 'N',
        '??': 'U',
        // lithuanian
        '??': 'e',
        '??': 'i',
        '??': 'u',
        '??': 'E',
        '??': 'I',
        '??': 'U',
        // romanian
        '??': 't',
        '??': 'T',
        '??': 't',
        '??': 'T',
        '??': 's',
        '??': 'S',
        '??': 'a',
        '??': 'A',
        // vietnamese
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'A',
        '???': 'E',
        '???': 'E',
        '???': 'E',
        '???': 'E',
        '???': 'E',
        '???': 'E',
        '???': 'E',
        '???': 'E',
        '???': 'I',
        '???': 'I',
        '??': 'I',
        '???': 'O',
        '???': 'O',
        '???': 'O',
        '???': 'O',
        '???': 'O',
        '???': 'O',
        '???': 'O',
        '??': 'O',
        '???': 'O',
        '???': 'O',
        '???': 'O',
        '???': 'O',
        '???': 'O',
        '???': 'U',
        '???': 'U',
        '??': 'U',
        '??': 'U',
        '???': 'U',
        '???': 'U',
        '???': 'U',
        '???': 'U',
        '???': 'U',
        '???': 'Y',
        '???': 'Y',
        '???': 'Y',
        '???': 'Y',
        '??': 'D',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'a',
        '???': 'e',
        '???': 'e',
        '???': 'e',
        '???': 'e',
        '???': 'e',
        '???': 'e',
        '???': 'e',
        '???': 'e',
        '???': 'i',
        '???': 'i',
        '??': 'i',
        '???': 'o',
        '???': 'o',
        '???': 'o',
        '???': 'o',
        '???': 'o',
        '???': 'o',
        '???': 'o',
        '??': 'o',
        '???': 'o',
        '???': 'o',
        '???': 'o',
        '???': 'o',
        '???': 'o',
        '???': 'u',
        '???': 'u',
        '??': 'u',
        '??': 'u',
        '???': 'u',
        '???': 'u',
        '???': 'u',
        '???': 'u',
        '???': 'u',
        '???': 'y',
        '???': 'y',
        '???': 'y',
        '???': 'y',
        '??': 'd',
        // currency
        '???': 'euro',
        '???': 'cruzeiro',
        '???': 'french franc',
        '??': 'pound',
        '???': 'lira',
        '???': 'mill',
        '???': 'naira',
        '???': 'peseta',
        '???': 'rupee',
        '???': 'won',
        '???': 'new shequel',
        '???': 'dong',
        '???': 'kip',
        '???': 'tugrik',
        '???': 'drachma',
        '???': 'penny',
        '???': 'peso',
        '???': 'guarani',
        '???': 'austral',
        '???': 'hryvnia',
        '???': 'cedi',
        '??': 'cent',
        '??': 'yen',
        '???': 'yuan',
        '???': 'yen',
        '???': 'rial',
        '???': 'ecu',
        '??': 'currency',
        '???': 'baht',
        '$': 'dollar',
        '???': 'indian rupee',
        // symbols
        '??': '(c)',
        '??': 'oe',
        '??': 'OE',
        '???': 'sum',
        '??': '(r)',
        '???': '+',
        '???': '"',
        '???': '"',
        '???': "'",
        '???': "'",
        '???': 'd',
        '??': 'f',
        '???': 'tm',
        '???': 'sm',
        '???': '...',
        '??': 'o',
        '??': 'o',
        '??': 'a',
        '???': '*',
        '???': 'delta',
        '???': 'infinity',
        '???': 'love',
        '&': 'and',
        '|': 'or',
        '<': 'less',
        '>': 'greater'
    };

    QorSlug.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {

                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorSlug(this, options)));
            }

            if (typeof options === 'string' && $.isFunction(fn = data[options])) {
                fn.apply(data);
            }
        });
    };


    $(function () {
        var selector = '[data-toggle="qor.slug"]';

        $(document).
        on(EVENT_DISABLE, function (e) {
            QorSlug.plugin.call($(selector, e.target), 'destroy');
        }).
        on(EVENT_ENABLE, function (e) {
            QorSlug.plugin.call($(selector, e.target));
        }).
        triggerHandler(EVENT_ENABLE);
    });

    return QorSlug;

});
