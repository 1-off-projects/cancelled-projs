from time import sleep as wait
from playsound import playsound as ps
import os
script_dir = os.path.dirname(os.path.abspath(__file__))
sound_file_name = 'beat.mp3'
sound_file_path = os.path.join(script_dir, sound_file_name)
bpm = (float((input('bpm:')))/1000)
if bpm < 40:
    print('too slow')
    wait(2.5)
    exit
elif bpm > 210:
    print('too fast')
    wait(2.5)
    exit
else:
    while True:
        ps(sound_file_path)
        wait(bpm)