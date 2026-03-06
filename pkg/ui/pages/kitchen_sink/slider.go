package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func SliderSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("max-w-sm"),
			Input(
				X.Type("range"),
				X.Class("input w-full"),
				X.Attr("min", "0"),
				X.Attr("max", "27"),
				X.Attr("value", "12"),
				X.Attr("style", "--slider-value: 44.44444444444444%;"),
			),
		),
		Raw(`
		<script>
		  (() => {
		    const sliders = document.querySelectorAll('input[type="range"].input');
		    if (!sliders) return;

		    const updateSlider = (el) => {
		      const min = parseFloat(el.min || 0);
		      const max = parseFloat(el.max || 100);
		      const value = parseFloat(el.value);
		      const percent = (max === min) ? 0 : ((value - min) / (max - min)) * 100;
		      el.style.setProperty('--slider-value', percent + '%');
		    };

		    sliders.forEach(slider => {
		      updateSlider(slider);
		      slider.addEventListener('input', (event) => updateSlider(event.target));
		    });
		  })();
		</script>
		`),
	)
}
