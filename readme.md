# scriptstick
![image of scriptstick](https://github.com/axellse/crackedbadusb/raw/refs/heads/main/media/poster.png)

The scriptstick is a versatile scripting device in the form of a usb flash drive. It is very similar to the traditional badusb, but since it uses Lua instead of DuckyScript, it can do so much more than act as a keyboard. The onboard hardware also allows for things like meshtastic. The firmware writen in go with tinygo.

## why?

I made this project because i love the form factor of a usb stick, and i felt that a lot of fun could be had by expanding the capabilities of traditional bad usbs to more than just HID. By adding just a couple of hardware components you can make a regular looking usb stick so incredily versatile. Pair this with being able to change function easily through a dip switch and you've got yourself a swiss army knife for computers. (holy phone salesman)

## how?
I have compiled a couple examples of what the scriptstick could do (theoretically, since i dont have the device yet):
- badusb (ofc)
- security key where the private keys are stored on an sd card which you could hide somewhere.
- boot a very small OS (in theory)
- IR LED to MOSFET pads + usb otg adapter + android phone + custom app = portable universal tv remote
- meshtastic!!! on the go maybe with an android phone?
- communicate with another scriptstick up to kilometers away!
- sd card to usb adapter (not the fastest)
- not sure if mosfets have any/enough capacitance to make this not work but you could build a pretty shitty mp3 player.
- modify a keyboard with a lora module inside, plug scriptstick into your pc, use your keyboard from kilometers away. (why)

The scriptstick project is sponsored by [Hack club](https://hack.club).

## blueprint approval stuff
### schematic
<img width="3507" height="2480" alt="bild" src="https://github.com/user-attachments/assets/782f1952-c9c8-4c0c-89cf-4bb6f828bd42" />

### pcb
<img width="1919" height="1031" alt="bild" src="https://github.com/user-attachments/assets/e3a9cba8-02fd-4a04-b5de-d51880dd02eb" />

### BOM
|Name                   |Quantity|Unit cost (SEK)|Variant                    |Manufacturer/seller    |Cost (SEK)|Link                       |Comment                                                  |
|-----------------------|--------|---------------|---------------------------|-----------------------|----------|---------------------------|---------------------------------------------------------|
|USB C to A adapter     |1       |22,17          |Blue A to C                |Elough Store/Aliexpress|22,17     |Link                       |Required as my PC only has USB A and not C               |
|Components             |3       |134,67         |N/A                        |LCSC                   |404,01    |N/A (Private shopping cart)|See production/lcsc_bom.csv in the git repo              |
|LCSC Shipping Fee      |N/A     |N/A            |Global Direct Standard Line|LCSC                   |91,75     |N/A                        |                                                         |
|PCBs                   |10      |4,73           |Black 2-layer              |JLCPCB                 |47,3      |N/A (Private shopping cart)|See production/gerber.zip in the git repo.               |
|PCB Stencil            |1       |108,42         |Top+Bottom on same stencil |JLCPCB                 |108,42    |N/A (Private shopping cart)|I assemble myself to minimize costs                      |
|JLCPCB Shipping Fee    |N/A     |N/A            |Global Direct Standard Line|JLCPCB                 |179,73    |N/A                        |                                                         |
|Aliexpress Shipping Fee|N/A     |N/A            |Choice shipping (cheapest) |Aliexpress             |18,68     |N/A                        |                                                         |

|Total (SEK)            |N/A|N/A   |N/A                        |N/A                    |872,06|N/A                        |Note that Swedish 25% sales tax is included in all prices|
|-----------------------|---|------|---------------------------|-----------------------|------|---------------------------|---------------------------------------------------------|
|Total (USD): 95$       |   |      |                           |                       |      |                           |                                                         |
