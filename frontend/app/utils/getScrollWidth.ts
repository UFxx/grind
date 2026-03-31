export default () =>
{
	setTimeout(() =>
	{
		document.body.style.setProperty('--scrollbar-width', (window.innerWidth - document.body.clientWidth) + 'px');
	}, 510)
}
