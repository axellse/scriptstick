# scriptstick
![image of scriptstick](https://github.com/axellse/crackedbadusb/raw/refs/heads/main/media/poster.png)

The scriptstick is a versatile scripting device in the form of a usb flash drive. It is very similar to the traditional badusb, but since it uses Lua instead of DuckyScript, it can do so much more than act as a keyboard. The onboard hardware also always for things like meshtastic. The firmware writen in tinygo.

## why?

I made this project because i love the form factor of a usb stick, and i felt that a lot of fun could be had by expanding a bad usb's capabilities to more than just HID. By adding just a couple of hardware components you can make a regular looking usb stick so incredily versatile. Pair this with being able to change function easily through a dip switch and you've got yourself a swiss army knife for computers. (holy phone salesman)

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
here is the schematic for the project:

<img width="3507" height="2480" alt="bild" src="https://github.com/user-attachments/assets/782f1952-c9c8-4c0c-89cf-4bb6f828bd42" />

here is the pcb inside of kicad:

<img width="1919" height="1031" alt="bild" src="https://github.com/user-attachments/assets/e3a9cba8-02fd-4a04-b5de-d51880dd02eb" />
