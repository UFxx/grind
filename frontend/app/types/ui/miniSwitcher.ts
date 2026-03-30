type UiMiniSwitcherId = 'ai' | 'manually';

export interface UiMiniSwitcher
{
	items      : [UiMiniSwitcherItem, UiMiniSwitcherItem]
	modelValue : UiMiniSwitcherItem
};

export interface UiMiniSwitcherItem
{
	id    : UiMiniSwitcherId,
	label : string,
}