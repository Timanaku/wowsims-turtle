import { BASE_PATH } from '../../constants/other';
import { ItemSlot } from '../../proto/common';

const emptySlotIcons: Record<ItemSlot, string> = {
	[ItemSlot.ItemSlotHead]: `${BASE_PATH}assets/item_slots/head.jpg`,
	[ItemSlot.ItemSlotNeck]: `${BASE_PATH}assets/item_slots/neck.jpg`,
	[ItemSlot.ItemSlotShoulder]: `${BASE_PATH}assets/item_slots/shoulders.jpg`,
	[ItemSlot.ItemSlotBack]: `${BASE_PATH}assets/item_slots/shirt.jpg`,
	[ItemSlot.ItemSlotChest]: `${BASE_PATH}assets/item_slots/chest.jpg`,
	[ItemSlot.ItemSlotWrist]: `${BASE_PATH}assets/item_slots/wrists.jpg`,
	[ItemSlot.ItemSlotHands]: `${BASE_PATH}assets/item_slots/hands.jpg`,
	[ItemSlot.ItemSlotWaist]: `${BASE_PATH}assets/item_slots/waist.jpg`,
	[ItemSlot.ItemSlotLegs]: `${BASE_PATH}assets/item_slots/legs.jpg`,
	[ItemSlot.ItemSlotFeet]: `${BASE_PATH}assets/item_slots/feet.jpg`,
	[ItemSlot.ItemSlotFinger1]: `${BASE_PATH}assets/item_slots/finger.jpg`,
	[ItemSlot.ItemSlotFinger2]: `${BASE_PATH}assets/item_slots/finger.jpg`,
	[ItemSlot.ItemSlotTrinket1]: `${BASE_PATH}assets/item_slots/trinket.jpg`,
	[ItemSlot.ItemSlotTrinket2]: `${BASE_PATH}assets/item_slots/trinket.jpg`,
	[ItemSlot.ItemSlotMainHand]: `${BASE_PATH}assets/item_slots/mainhand.jpg`,
	[ItemSlot.ItemSlotOffHand]: `${BASE_PATH}assets/item_slots/offhand.jpg`,
	[ItemSlot.ItemSlotRanged]: `${BASE_PATH}assets/item_slots/ranged.jpg`,
};
export function getEmptySlotIconUrl(slot: ItemSlot): string {
	return emptySlotIcons[slot];
}
