  
		AACHIP AAX52IDE V2.0中文使用說明書

8051 8052 皆為intel 註冊的商標及專利.本公司所寫的AACHIPX52IDE,
及X52 ASSEMBLY COMPILER 皆為本公司參考網路技術手冊所寫,與INTEL沒有關係,此程式為本公司內部使用,
及社團個人學習討論使用,特此聲明.如有疑慮請勿使用.

1.先下載AACHIPX52IDE.exe
2.執行AACHIPX52IDE.exe
3.會先顯示是不是要在視窗上顯示捷徑,請打勾.
4.會安裝在c:\AACHIP\  請勿更改路徑
5.安裝好後在C:\AACHIP\裡面會有AACHIP.ico,aachipx52.ini,AACHIPX52IDE.exe,unins000.dat,unins000.exe,
  AACX52IDE使用說明書.txt這6個檔案請勿刪除.
6.安裝好後會自動啟動,一般在3秒內會啟動成功,但是32bit cpu啟動可能會啟動成功需要到25秒左右.
7.本ide支援,SSFDC C.Compiler 及 本公司自行研發aax52 Assemblyer 組譯	軟體.
8.本ide 沒有文字編輯器及vcd wave 觀察器請自行從網路下載

9.文字編輯器可以使用notepad++,visual sutudio code,ultraedit 等文字編輯器
	文字編輯器主要是用來編輯8051 assembly 程式語言,和sdcc C語言編輯.
	notepad++ https://notepad-plus-plus.org/downloads/
	visual sutudio code https://code.visualstudio.com/Download

10.vcd wave 觀察器,可以使用gtkwave(免費),modelsim(學生版)
   vcd wave 觀察器 主要是查看整個8051 io port 輸出的hi lo情況 方便查看波形
	gtkwave  https://sourceforge.net/projects/gtkwave/files/gtkwave-3.3.100-bin-win64
	modelsim https://eda.sw.siemens.com/en-US/ic/modelsim/

11先用文字編輯器打開c:\AACHIP\aachipx52.ini
目前只有兩行
	TEXTEDIT_PATH=C:\\Program Files\\Notepad++\\notepad++.exe
	WAVEEDIT_PATH=C:\gtkwave-3.3.90-bin-win64\gtkwave64\bin\gtkwave.exe
把你下載執行的 編輯程式 和 vcd觀看程式.exe所在的位置填入"=右邊" ;  "=左邊請勿改變更動"
PS.如果你有使用本公司USB MULITY IO 只要連上軟體後就會自動產生LICENSE KEY為第3行

12.先介紹AACHIPX52IDE專用的ASSEMBLY 組譯程式.
12_1.先寫一個簡單的8051 Assrmbly 程式,(如果沒有下載文字編輯器,不能使用此功能)
a.利用x52組合語言建置
b.點選x52組合語言建置,點選建置新ASSEMBLY
c.寫一個短程式
//----------------------------------------
//這是第一個測試的程式
//不使用任何inc 或是h的附加lib
//測試p0~p3 output
.org 0h
	ljmp main
	
.org 0150h
main:
    //先讓port0,port1,port2,port3都為lo
	mov	a,#00
	mov	p0,a
	mov	p1,a
	mov	p2,a
	mov	p3,a
	
	mov r0,#8
	mov a,#01
    loop0:
        mov P0,a; //讓p0=a
        lcall delay ; //delay 5
		rl a;			//左移一位
		DJNZ R0, loop0  //如果r0不等於0則跳至loop1
		
	mov r0,#8
	mov a,#01
    loop1:
        mov P1,a; //讓p1=a
        lcall delay ; //delay 5
		rl a;			//左移一位
		DJNZ R0, loop1  //如果r0不等於0則跳至loop1	

	mov r0,#8
	mov a,#01
    loop2:
        mov P2,a; //讓p2=a
        lcall delay ; //delay 5
		rl a;			//左移一位
		DJNZ R0, loop2  //如果r0不等於0則跳至loop2	
		
	mov r0,#8
	mov a,#01
    loop3:
        mov P3,a; //讓p3=a
        lcall delay ; //delay 5
		rl a;			//左移一位
		DJNZ R0, loop3  //如果r0不等於0則跳至loop3	
		
	loop4:
		ljmp	loop4	//做完後停在這裡
		
		
        
delay:
    mov r1,#5
delay_loop:
    djnz r1,delay_loop
    ret
//======================================


//--------------------------------------------
d.寫好後一定要存檔案,例如存成test1.asm(建置副屬檔要.asm 因為此ide會讀取.asm這個檔案,如果沒有存,則無法compiler )
e.重新用 點選 x52組合語言建置->開啟舊檔,將剛剛編輯好的test1.asm 叫進來
e.點選 x52組合語言建置->點選 COMPILER建置
f.如果沒有問題則output視窗會出現pass1 ,pass2 則compiler ok 你在test1.asm的目錄裡面會出現test1.dis,test1.lst,test1.hex,test1.bin

13.進入模擬器
a.點選 檔案->打開AAX52偵錯bin+DIS檔案,選擇剛剛你的目錄下會有一個test.bin
b.這時候在Deassembly Debug 會出現反組譯的程式碼(系統會順便將test.dis 放到Deassembly Debug);
	右邊框的Memory 位置00000000一開始應該會有cpu reset碼,如果00000000位置為00 00 00,則有可能是你設定org,
	如果有強制設定org則一定要一開始設org $00 ,再重新compiler
c.一切沒有問題後就可以開始偵錯.
e.偵錯主要功能有RESET,RUN,STEPIN,STEPOUT,STOP
RESET:讓CPU 回到開始.
RUN:讓CPU 開始執行.
STEPIN:讓CPU 單步執行後停止.
STEPOUT:讓CPU 單步執行後停止,但是遇到LCALL,ACALL會直接執行完後才停止,除非CALL中間有被設定中斷.
STOP:讓CPU停在這裡.

14.中斷設定,本ide提供3個中斷.
a.你可以選擇到你要中斷的行,滑鼠左鍵點一下,當行會先顯示藍色,再按keyboard "F9"此時當行會變成紅色,
	這時候這行就是中斷點.如果再按一次"F9"則又會變藍色,取消中斷點.

15.程式執行紀錄(目前只有支援Simulation)
a.開始記錄:在任何時間你都可以啟動開始記錄,但是必須在cpu停止時啟動,
	啟動完後就會開始記錄整個cpu的內部暫存器及pc所有執行的過程,方便你除錯.
b.停止紀錄:停止紀錄正在紀錄的資料.
c.儲存CODE執行紀錄:將整個執行的所有資料存成檔案,方便你除錯.

16.IO記錄分析(目前只有支援Simulation):用來分析你的IOPORT 輸出的時間是不是和你想要的時間一樣,
	這個功能一般是用在IC設計,時用來查看 IO 和 時間 對應的關連. 加強偵測.
a.開始記錄:在任何時間你都可以啟動開始記錄,但是必須在cpu停止時啟動,啟動完後就會開始記錄整個cpu的IO PORT.
b.關閉紀錄:在任何時間你都可以啟動關閉記錄,但是必須在cpu停止時關閉,關閉後就可以停止紀錄.
c.存成vcd檔案:把你剛剛記錄的io port 轉成vcd,並且存成檔案
d.打開VCD分析工具:如果你有完成前面3個步驟,就可以查看整個IO PORT 和 時間的關聯.

17.連接硬體:FPGA 連接 USB DEBUG,FPGA UPDATE
a.FPGA 連接 USB DEBUG->本公司利用ftdi ft2232h設計一款usb type c 自動判斷外部實驗版的工作電壓來做jtag 及debug mode 來控制
	我們公司另一款FPGA AACHIP Emulator v2.0 可以完美模仿8052的功能.
b.FPGA UPDATE 更新FPGA 8052 功能專用.提供客戶自行更新FPGA 新功能.
	  

18.支援ssfdc C COMPILER 出來的ihx和map,本AAX52IDE會自動將ihx轉換成disassembly 再和map,lst就可以結合c source file 來做比對,
	方便你快速寫c程式再來simulator

19.顯示:各種功能視窗顯示.
a.主畫面記憶體顯示:當你消除視窗上面的功能可以再利用這裡把原來的視窗叫回來
a1:顯示sfr_reg.
a2:顯示8051cpu內部ram.
a3:顯示extram和rom code
a4:顯示正在執行cpu 內部的reg
b.除錯顯示:Deassembly Debug 和 output 消除後可以再叫回來
c:顯示8組 io port :目前支援port0,port1,port2,port3,這4個PORT 相容於8052,另外有3組獨立的8BIT IO PORT及一組獨立的8BIT OUTPUT PORT
d:顯示timer:目前支援timer0,timer1,timer2, TIMER2 支援DEC 
e:顯示所有設定的中斷:目前支援3個中斷點,每次設定好一個中斷點,這裡就會顯示
註:中斷只能在主畫面的Deassembly Debug上面設定,不能在獨立的Deassembly Debug上面設定.
f:啟動所有設定的中斷:把所有的中斷都致能,再按一次就會停止所有的中斷.
g:各種獨立顯示的視窗顯示:可以獨立於主畫面之外顯示各種功能視窗.

20.V20版新增加Cwindows 可以在開啟sdcc C COMPILER 出來的ihx和map,LST,結合C SOURCE FILE 就可以直接看C SOURCE 來做Debug
  1.點選Cwindows 的C source file 的行,在debug disassembly file 就會指向對應的disassembly 指令的行
  2.當程式開始執行後無論是單步還是執行都可以完美對應到C Source file.
  
********詳細說明可以參考 www.youtube.com/@AACHIP-ku3kt 內有詳細說明及例子**********************************************
