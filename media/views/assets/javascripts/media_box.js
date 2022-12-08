"use strict";function _typeof(e){return(_typeof="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e})(e)}!function(e){"function"==typeof define&&define.amd?define(["jquery"],e):"object"===("undefined"==typeof exports?"undefined":_typeof(exports))?e(require("jquery")):e(jQuery)}(function(g){var o=g("body"),e=g(document),n="qor.medialibrary.select",t="click."+n,i="enable."+n,a="reload.qor.bottomsheets",s=".qor-select__select-icon",c=".qor-field__mediabox-list",y=".qor-field__mediabox-item",d=".qor-field__mediabox-data",r="is_selected",l="is_deleted",v="textarea.qor-file__options",m="qor-bottomsheets__mediabox";function h(e,t){this.$element=g(e),this.options=g.extend({},h.DEFAULTS,g.isPlainObject(t)&&t),this.init()}return h.prototype={constructor:h,init:function(){var e=this.$element;this.SELECT_MEDIABOX_UNDO_TEMPLATE=e.find('[name="media-box-undo-delete"]').html(),this.bind(),this.initSelectedMedia()},bind:function(){e.off(a).on(a,".".concat(m),this.reloadData.bind(this)),this.$element.off(t).off("change.qor.cropper").on(t,"[data-mediabox-url]",this.openBottomSheets.bind(this)).on(t,".qor-cropper__toggle--delete",this.deleteSelected.bind(this)).on(t,".qor-cropper__toggle--undo",this.undoDeleteSelected.bind(this)).on("change.qor.cropper",v,this.imageCrop.bind(this))},unbind:function(){e.off(a,".".concat(m)),this.$element.off(t).off("change.qor.cropper")},deleteSelected:function(e){var t=g(e.target);return t.closest(y).addClass(l).append(this.SELECT_MEDIABOX_UNDO_TEMPLATE).find(".qor-file__list").hide(),this.updateMediaLibraryData(t.closest(c)),this.$element.find(d).data("isDeleted",!0),g("body").trigger("itemRemoved.medialibrary",this),!1},undoDeleteSelected:function(e){var t=g(e.target);return t.closest(y).removeClass(l).find(".qor-file__list").show(),this.updateMediaLibraryData(t.closest(c)),t.closest(".qor-fieldset__alert").remove(),this.$element.find(d).data("isDeleted",!1),!1},imageCrop:function(e){var t=g(e.target).closest(y);this.syncImageCrop(t,this.resetImages)},openBottomSheets:function(e){var t,i=g(e.target).closest("[data-mediabox-url]"),a=i.data();a.isDisabled||(this.BottomSheets=o.data("qor.bottomsheets"),this.bottomsheetsData=a,this.$parent=t=i.closest(".qor-field__mediabox"),this.$selectFeild=t.find(c),a.url=a.mediaboxUrl,this.SELECT_MANY_SELECTED_ICON=g('[name="media-box-select-many-selected-icon"]').html(),this.SELECT_MANY_HINT=g('[name="media-box-select-many-hint"]').html(),this.TEMPLATE_IMAGE=t.find('[name="media-box-template"]').html(),this.TEMPLATE_FILE=t.find('[name="media-box-file-template"]').html(),this.TEMPLATE_UPLOADEDVIDEO=t.find('[name="media-box-uploadedvideo-template"]').html(),this.TEMPLATE_VIDEOLINK=t.find('[name="media-box-videolink-template"]').html(),this.SELECT_MEDIABOX_UNDO_TEMPLATE=t.find('[name="media-box-undo-delete"]').html(),this.BottomSheets.open(a,this.handleSelectMany.bind(this)))},initSelectedMedia:function(){var e,t=this.$element,i=t.find(y),a=JSON.parse(t.find(d).val());if(a)for(var o=0;o<a.length;o++)(e=i.filter('[data-primary-key="'+a[o].ID+'"]')).data().description||e.data("description",a[o].Description)},initMedia:function(){var e,t,i,a=this.$selectFeild.find(y).not("."+l),o=this.$bottomsheets.find("tbody tr"),n=this;a.each(function(){i=g(this).data().primaryKey,e=o.filter('[data-primary-key="'+i+'"]').addClass(r),n.changeIcon(e,!0)}),o.each(function(){e=g(this),t=e.find(".qor-table--ml-slideout p img").first(),e.find(".qor-table__actions").remove(),t.length&&(e.find(".qor-table--medialibrary-item").css("background-image","url("+t.prop("src")+")"),t.parent().remove())}),"1"!=this.bottomsheetsData.maxItem&&this.updateHint(this.getSelectedItemData())},reloadData:function(){this.$selectFeild&&this.initMedia()},renderHint:function(e){return window.Mustache.render(this.SELECT_MANY_HINT,e)},getSelectedItemData:function(e){var t=(e||this.$selectFeild).find(y).not("."+l),i=[];return t.length&&t.each(function(){var e=g(this).data();i.push({ID:e.primaryKey,Url:e.originalUrl.replace(/.original.(\w+)$/,".$1"),Description:e.description,FileName:e.fileName,VideoLink:e.videolink})}),{files:i,selectedNum:i.length}},updateHint:function(e){var t;g.extend(e,this.bottomsheetsData),t=this.renderHint(e),g(".qor-selectmany__hint").remove(),this.$bottomsheets.find(".qor-page__body").before(t)},updateMediaLibraryData:function(e,t){var i=e?e.find(d):this.$selectFeild.find(d),a=this.getSelectedItemData(e);i.val(JSON.stringify(a.files)).data("mediaData",t).trigger("changed.medialibrary",[t])},changeIcon:function(e,t){var i=e.find(".qor-table--medialibrary-item"),a=i.length?i:e.find("td:first");e.find(s).remove(),t&&("one"==t&&g("."+m).find(s).remove(),a.prepend(this.SELECT_MANY_SELECTED_ICON))},syncImageCrop:function(t,i){var a,o,n=JSON.parse(t.find(v).val()),e=t.closest("form").find(":submit"),s=t.data().mediaLibraryUrl,d=this,r={},l=["Width","Height"],m=t.find("img[data-size-name]");delete n.ID,delete n.Url,n.Sizes={},m.each(function(){if((o=g(this).data()).sizeResolutionWidth||o.sizeResolution){n.Sizes[o.sizeName]={};for(var e=0;e<l.length;e++)a=(a=o["sizeResolution"+l[e]])||o.sizeResolution[l[e]],n.Sizes[o.sizeName][l[e]]=a}}),r.MediaOption=JSON.stringify(n),this.addMediaLoading(t),g.ajax({type:"PUT",url:s,data:JSON.stringify(r),contentType:"application/json",dataType:"json",beforeSend:function(){e.attr("disabled",!0)},success:function(e){r.MediaOption=JSON.parse(e.MediaOption),t.data().originalUrl=r.MediaOption.OriginalURL,d.updateMediaLibraryData(t.closest(c),r),i&&g.isFunction(i)&&i(r,t)},complete:function(){e.attr("disabled",!1),t.find(".qor-media-loading").remove()}})},showHiddenItem:function(e){e.removeClass(l).find(".qor-file__list").show(),e.find(".qor-fieldset__alert").remove()},removeItem:function(e){var t=e.primaryKey;this.$selectFeild.find('[data-primary-key="'+t+'"]').remove(),this.changeIcon(e.$clickElement)},compareCropSizes:function(e){var t,i,a=e.MediaOption.CropOptions,o=this.bottomsheetsData.cropSizes;if(!o||"image"!=e.SelectedType)return!1;if(void 0===a)return!0;if(t=(o=o.split(",")).length-1,!window._.isObject(a))return!1;if((i=Object.keys(a)).length)for(var n=0;n<t;n++)if(-1==i.indexOf(o[n]))return!0;return!1},addItem:function(e,t){var i,a,o,n=g(window.Mustache.render(this.TEMPLATE_IMAGE,e)),s=n.find(".qor-file__input"),d=s.closest(y),r=d.closest("form").find(":submit"),l=this.$selectFeild.find('[data-primary-key="'+e.primaryKey+'"]'),m=this.bottomsheetsData.maxItem,c=this.getSelectedItemData().selectedNum,h=e.MediaOption.CropOptions,f=this.compareCropSizes(e),u=e.SelectedType,p=/.svg$/.test(e.MediaOption.FileName),b=this;if(r.attr("disabled",!0),t||(1==m?this.changeIcon(e.$clickElement,"one"):this.changeIcon(e.$clickElement,!0)),m&&m<=c){if(1!=m)return void window.alert(this.bottomsheetsData.maxItemHint);this.$selectFeild.find(y).remove()}if(l.length)return this.showHiddenItem(l),void(1==m&&setTimeout(function(){b.$bottomsheets.remove(),g(".qor-bottomsheets").is(":visible")||g("body").removeClass("qor-bottomsheets-open")},1e3));1==m&&this.$selectFeild.find(y).filter(".is_deleted").remove(),p||("video"===u?n=g(window.Mustache.render(this.TEMPLATE_UPLOADEDVIDEO,e)):"video_link"===u?(a=e.MediaOption.Video,i=!(!(o=a.match(/^.*((youtu.be\/)|(v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))\??v?=?([^#\&\?]*).*/))||11!=o[7].length)&&o[7],e.VideoLink=i?"//www.youtube.com/embed/".concat(i,"?rel=0&fs=0&modestbranding=1&disablekb=1"):e.MediaOption.Video,n=g(window.Mustache.render(this.TEMPLATE_VIDEOLINK,e))):"file"===u&&(n=g(window.Mustache.render(this.TEMPLATE_FILE,e)))),n.data({description:e.MediaOption.Description,mediaData:e,videolink:"video_link"===u?e.MediaOption.Video:""}),p&&n.addClass("is-svg").find(".qor-file__input").remove(),n.appendTo(this.$selectFeild),h&&"image"===u&&this.resetImages(e,n),"image"===u&&n.find(v).val(JSON.stringify(e.MediaOption)),n.trigger("enable"),f&&s.data("qor.cropper")&&!p&&s.data("qor.cropper").load(e.MediaOption.URL,!0,function(){b.syncImageCrop(d,b.resetImages)}),this.$bottomsheets.find(".qor-media-loading").remove(),g("body").trigger("itemAdded.medialibrary",[this,e]),!t&&1!=m||setTimeout(function(){b.$bottomsheets.remove(),g(".qor-bottomsheets").is(":visible")||g("body").removeClass("qor-bottomsheets-open")},150)},resetImages:function(e,t){var a=e.MediaOption.CropOptions,i=Object.keys(a),o=e.MediaOption.OriginalURL;if(/\.original\./.test(o)){for(var n=i.length-1;0<=n;n--)a[i[n]].URL=o.replace(/\.original\./,".".concat(i[n],"."));t.find("img[data-size-name]").each(function(){var e=g(this),t=(Math.random()+1).toString(36).substring(7),i=e.data().sizeName;"original"!=i&&a[i]&&e.prop("src","".concat(a[i].URL,"?").concat(t))})}},handleSelectMany:function(e){var t={loading:this.addMediaLoading,onSelect:this.onSelectResults.bind(this),onSubmit:this.onSubmitResults.bind(this)};e.qorSelectCore(t).addClass(m),this.$bottomsheets=e,this.initMedia()},onSelectResults:function(e){this.handleResultsData(e)},onSubmitResults:function(e){this.handleResultsData(e,!0)},handleResultsData:function(e,t){var i=e.$clickElement;if(e.mediaLibraryUrl||t||(e.mediaLibraryUrl=e.url),t)return e.mediaLibraryUrl=this.bottomsheetsData.mediaboxUrl.split("?")[0]+"/"+e.primaryKey,this.addItem(e,t),void this.updateDatas(e);i.toggleClass(r),i.hasClass(r)?this.addItem(e):this.removeItem(e),this.updateDatas(e)},addMediaLoading:function(e){g('<div class="qor-media-loading"><div class="mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active"></div></div>').appendTo(e).trigger("enable.qor.material")},updateDatas:function(e){"1"!=this.bottomsheetsData.maxItem&&this.updateHint(this.getSelectedItemData()),this.updateMediaLibraryData(null,e)},destroy:function(){this.unbind(),this.$element.removeData(n)}},h.plugin=function(a){return this.each(function(){var e,t=g(this),i=t.data(n);if(!i){if(/destroy/.test(a))return;t.data(n,i=new h(this,a))}"string"==typeof a&&g.isFunction(e=i[a])&&e.apply(i)})},g(function(){var t='[data-toggle="qor.mediabox"]';g(document).on("disable.qor.medialibrary.select",function(e){h.plugin.call(g(t,e.target),"destroy")}).on(i,function(e){h.plugin.call(g(t,e.target))}).triggerHandler(i)}),h});