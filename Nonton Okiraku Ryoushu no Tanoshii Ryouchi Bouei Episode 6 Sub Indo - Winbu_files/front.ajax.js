! function (a) {
	a(function () {
		var d = "#east_option";
		var c, f;

		function b() {
			var e = "#countdown";
			if (1 == a(e).length) {
				var g;
				if (g = a(e).text(), 1 == (g = parseInt(g, 10))) {
					return void a(e).remove()
				}
				g--, a(e).html(g), setTimeout(b, 1000)
			}
		}
		a(document).ready(function () {
			a(document).on("click", ".process_list", function () {
				var e = a(this);
				return post_id = e.attr("data-post-id"), security = e.attr("data-nonce"), a(".fico").removeClass("fa-heart"), a(".fico").addClass("fa-spinner loading"), a.ajax({
					type: "POST",
					dataType: "json",
					url: easthemeajax.url,
					data: {
						action: "east_process_favorites",
						post_id: post_id,
						nonce: security
					},
					error: function (g) {},
					success: function (g) {
						a(".fico").addClass("fa-heart"), a(".fico").removeClass("fa-spinner loading"), a(".process_list").toggleClass("in-list"), a(".textfavs").html(g.text), a(".countpep").html(g.count)
					}
				}), !1
			}), a(document).on("click", ".remove_favorites", function () {
				var e = a(this);
				return post_id = e.attr("data-postid"), security = e.attr("data-nonce"), a("#post" + post_id).remove(), a.ajax({
					type: "POST",
					url: easthemeajax.url,
					data: {
						action: "east_process_favorites",
						post_id: post_id,
						nonce: security,
						total: "on"
					},
					success: function (g) {
						a(".totalfavorites_user").html(g)
					}
				}), !1
			}), a(document).on("click", ".user_views_control", function () {
				var e = a(this);
				return post_id = e.attr("data-postid"), security = e.attr("data-nonce"), a("#watch" + post_id).remove(), a.ajax({
					type: "POST",
					url: easthemeajax.url,
					data: {
						action: "east_proccess_watch",
						post_id: post_id,
						nonce: security,
						total: "on"
					},
					success: function (g) {
						a(".totalviews_user").html(g)
					}
				}), !1
			}), a(document).on("click", ".process_views", function () {
				var e = a(this);
				return post_id = e.attr("data-post-id"), security = e.attr("data-nonce"), a(".wvcico").removeClass("fa-clock-o"), a(".wvcico").addClass("fa-spinner loading"), a.ajax({
					type: "POST",
					url: easthemeajax.url,
					data: {
						action: "east_proccess_watch",
						post_id: post_id,
						nonce: security
					},
					error: function (g) {},
					success: function (g) {
						a(".wvcico").addClass("fa-clock-o"), a(".wvcico").removeClass("fa-spinner loading"), a(".process_views").toggleClass("in-views"), a(".views-count-" + post_id).html(g)
					}
				}), !1
			}), a(".east_player_option").click(function () {
				var e = a(this).data("post");
				return nume = a(this).data("nume"), type = a(this).data("type"), tviw = a(".playerload").data("text"), scds = easthemeajax.playeradstime, a("#fakeplayer").hide(), a(this).hasClass("on") || (a(".east_player_option").removeClass("on"), a("#player-option-" + nume + " > span.loader").show(), a("#playeroptions").addClass("onload"), 0 < scds && (a(this).addClass("on")), a.ajax({
					type: "POST",
					url: easthemeajax.url,
					data: {
						action: "player_ajax",
						post: e,
						nume: nume,
						type: type
					},
					error: function (g) {},
					success: function (g) {
						0 < scds && b(), a("#player-option-" + nume + " > span.loader").hide(), a(".movieplay").html('<div class="preplayer"></div>'), setTimeout(function () {
							a(".asgdc").hide(), a(".movieplay").html('<div class="pframe">' + g + "</div>"), a("#playeroptions").removeClass("onload"), 0 < scds && (a(".playerload").html(tviw))
						}, 1000 * scds)
					}
				})), !1
			});
            1 == a("#player-option-1").length && 1 == easthemeajax.autoplayer && setTimeout(function () {
                a("#player-option-1").trigger("click");
			}, (c = 100, f = 700, Math.floor(Math.random() * (f - c + 1) + c))), 1 == a(d).length && (a(d).repeater({
				defaultValues: {
					type: easthemeajax.ltipe
				},
				show: function () {
					a(this).fadeIn("fast")
				},
				hide: function (e) {
					a(this).fadeOut("fast", function () {
						a(this).slideUp(e)
					})
				}
			}))
		})
	})
}(jQuery);
