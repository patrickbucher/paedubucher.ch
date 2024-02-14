---
title: Jahresrückblick 2021
subtitle: Gar nicht so schlecht?
author: Patrick Bucher
date: 2021-12-17T13:00:00+0100
lang: de
---

Selten war meine Laune am Jahresende so schlecht. Dass ich mitte Dezember
erschöpft bin und mich nach einer Pause sehne, ist nichts Neues und war zu
erwarten. Leider habe ich aber auch das Gefühl, dass dieses Jahr für mich extrem
unproduktiv war, und ich stehenbleibe, ja gar Rückschritte mache ‒ in allen
möglichen Lebensbereichen.

Ich möchte hier auf mein Jahr zurückblicken, soweit es mein Erinnerungsvermögen
erlaubt, um diesen Eindruck zu prüfen.

Was Fremdsprachen und Literatur betrifft, ist der Rückblick schnell abgehandelt.
Französisch und Russisch habe ich kaum verwendet; Englisch höchstens für die
Informatik.

An Literatur kann ich mich dieses Jahr kaum erinnern. Ayn Rands _Anthem_ habe
ich auf dem Kindle gelesen. In den Sommerferien in Arosa habe ich Michel
Houellebecqs _Unterwerfung_ aus einer Bibliotheksbox gezogen und dann in dieser
Woche gleich durchgelesen, ich musste es ja schliesslich vor meiner Abreise
wieder zurücklegen. Gontscharews _Oblomow_, den ich im Gepäck hatte, blieb dafür
ungelesen. Viele andere Romane habe ich _an-_ aber nicht _durchgelesen_.  Den
Kindle habe ich schliesslich wieder in den Keller verbannt, da ich damit nur
oberflächlich lese.

Was meinen Beruf und die Informatik angeht, weiss ich doch etwas mehr zu
berichten.

# Motiviert mit MOOCs

Ich weiss nicht, ob es Ende 2020 oder Anfang 2021 war, als ich den Kurs [Machine
Learning](https://www.coursera.org/learn/machine-learning?) abgeschlossen habe,
aber mit der [Machine Learning
Specialization](https://www.coursera.org/specializations/deep-learning) habe ich
wohl erst 2021 angefangen. Hier bin ich schnell vorangekommen und war lange Zeit
motiviert ‒ bis es schliesslich in einem Mooc zum Framework
[TensorFlow](https://www.tensorflow.org/) kam, das immer noch in der Version 1.x
verwendet worden war. Der Umstieg auf TensorFlow 2.x muss wohl ein grösseres
Problem sein, zumindest ist er mir nicht einmal für die Kursbeispiele gelungen.

Immerhin habe ich selber noch einige Beispiele entwickelt:
[x-o-classifier](https://github.com/patrickbucher/x-o-classifier) zur optischen
Unterscheidung der Zeichen `x` und `o`,
[digit-detection](https://github.com/patrickbucher/digit-detection) als
erweitertes OCR-Beispiel für die Ziffern von 0 bis 9, die [Titanic
Challenge](https://github.com/patrickbucher/titanic-challenge) und die [House
Prices](https://github.com/patrickbucher/house-prices) von Kaggle (wenn auch
mit mässigem Erfolg). Weiter sind mir kleinere Beispiele und eine
Zusammenfassung des ersten Kurses in meinem
[machine-learning-Repository](https://github.com/patrickbucher/machine-learning)
geblieben.

Für mich war das Thema Machine Learning dann vorerst erledigt. Ein vorbestelltes
und endlich im November eingetroffenes Buch ‒ [Math for Deep
Learning](https://nostarch.com/math-deep-learning) ‒ erinnerte mich an meine
verflogene Motivation.

# Funktionale Programmierung

Weiter ging es mit der funktionalen Programmierung, wozu ich den recht
gelungenen
[Scala-Kurs](https://www.coursera.org/learn/scala-functional-programming) von
Martin Odersky durchgearbeitet habe. Den meiner Meinung nach am besten
aufgebauten Kurs zur funktionalen Programmierung, [Programming Languages, Part
A](https://www.coursera.org/learn/programming-languages), konnte ich leider
nicht abschliessen: das ständige Lernen von Videos hat mich zu sehr erschöpft,
und ich wollte in meiner Freizeit weg vom Bildschirm.

So machte ich mich an
[SICP](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html)
heran, worin ich leider im ersten Kapitel steckengeblieben bin. Immerhin habe
ich die Übungen bis dahin [seriös
gelöst](https://github.com/patrickbucher/sicp/tree/master/ch01), sodass ich hier
bei späterer Gelegenheit wieder einsteigen kann.

Eigentlich wollte ich mich 2021 mit Erlang befassen, bin aber dann aus
pragmatischen Gründen doch auf [Elixir](https://elixir-lang.org/) eingeschwenkt.
Ich versuchte das Buch [Learn Functional Programming with
Elixir](https://pragprog.com/titles/cdc-elixir/learn-functional-programming-with-elixir/)
durchzuarbeiten, gab aber auf Seite 129 (von 198) auf, wovon der Kommentar ganz
unten in meinen
[Notizen](https://github.com/patrickbucher/learning-elixir/blob/master/learn-functional-programming-with-elixir/notes.md)
zeugt. Die Beispiele waren mir wohl einfach zu kindisch und praxisfern. Ich
versuchte es dann noch mit [Elixir in
Action](https://www.manning.com/books/elixir-in-action-second-edition), doch
hatte ich im Frühling dann andere Prioritäten.

Im Sommer arbeitete ich noch ein kleines eBook zur funktionalen Programmierung
in Python durch: [Functional Programming in
Python](https://leanpub.com/functionalprogramminginpython), wozu ich auch gleich
meine [persönliche
Zusammenfassung](https://github.com/patrickbucher/docs/blob/master/python/funcprog.md)
mit eigenen Beispielen schrieb. Einiges davon konnte ich dann gewinnbringend im
Arbeitsalltag einsetzen; Monaden gehören jedoch nicht dazu…

# Arbeit

In der Firma habe ich mich die ersten vier Monate v.a. mit der Migration von
Servern von Ubuntu 16.04 auf 20.04 befasst. U.a. musste ich einen Mailserver und
den ganzen Monitoring-Stack migrieren, was wohl das Schwierigste an der Aufgabe
war.

Für die Ingestion der Metriken via [Riemann](http://riemann.io/) musste ich, um
von Python 2 wegzukommen, von einer bestehenden Datenbank (deren Name ich
bereits vergessen habe) auf [InfluxDB](https://www.influxdata.com/) wechseln.
Hierzu durfte ich sogar etwas Clojure lernen, wobei mir [Getting
Clojure](https://pragprog.com/titles/roclojure/getting-clojure/) von Russ Olsen
geholfen hat. Ich benötigte nur ca. zwei Kapitel. Im Frühling arbeitete ich das
Buch dann in meiner Freizeit komplett durch, wovon mein Repository
[learning-clojure](https://github.com/patrickbucher/learning-clojure/) zeugt.

Insgesamt migrierte ich ca. 20 Server mit verschiedensten Anwendungen und
Technologien. Von Mailservern verstehe ich praktisch immer noch so wenig wie
zuvor; da es sich hier aber um etablierte Software handelt (Sendmail, Dovecot),
war die Umstellung nicht allzu schwierig und auch ohne tieferes Wissen zu
bewältigen.

Lehrreich war auch die Ablösung von [Bitbucket](https://bitbucket.org/product),
das ja seit Herbst nur noch in der Cloud und nicht mehr _on-premise_ läuft,
durch [Gitea](https://gitea.io/en-us/) ‒ eine deutlich schlankere Lösung, die in
[Go](https://go.dev) geschrieben ist.

Das wichtigste Projekt war jedoch der
[Flugzeugkonfigurator](https://www.seantis.ch/success-stories/aircraft-configurator-pilatus/),
mit dem man neu auch die Innenausstattung von Flugzeugen (PC-12 und PC-24)
konfigurieren kann. Obwohl mir die Programmierung grafischer Oberflächen
überhaupt nicht liegt, konnte ich auf Basis erstklassiger Renderings, die mir
vom Kunden geliefert worden sind, doch ein recht ansprechendes Ergebnis
erzielen.

Das Projekt hat mich ab Mai praktisch bis Ende Jahr beschäftigt und war äusserst
vielfältig: Neben Parsern in [PLY](https://www.dabeaz.com/ply/) ‒ und
entsprechenden Interpretern, wobei mein bescheidenes Wissen in LISP doch recht
hilfreich war ‒ habe ich viel Test-Driven Development mit
[PyTest](https://docs.pytest.org/) und einiges an Verarbeitung von Excel-Dateien
mit [OpenPyXL](https://openpyxl.readthedocs.io/en/stable/) gemacht. Auch die
PDF-Generierung mit [ReportLab](https://www.reportlab.com/opensource/) spielte
dabei eine wichtige Rolle, wobei solche Aufgaben wegen der manuellen Testerei
und Feinjustierung oft zermürbend sind. Viele Konstrukte erforderten rekursive
Verarbeitung; solche Programmieraufgaben machen mir immer am meisten Freude.

Den Frontend-Code, wozu ich Ende letzten Jahres v.a. JavaScript verwendet hatte,
konnte ich dieses Jahr reduzieren, indem ich viel Funktionalität ins Backend
verschob. Das war ein Erfolgserlebnis, zumal die Codebasis heute viel besser
aussieht als vor einem Jahr, obwohl sie geschätzt um die Hälfte angewachsen ist.
Auch die Testabdeckung liegt mittlerweile bei über 90%, wobei ich den
JavaScript-Code im Frontend nicht mitzähle. Hier habe ich leider noch keine
vernünftige und einfache Möglichkeit gefunden, den Code automatisch zu testen.

# Lehrlingsausbildung und Berufsschule

Seit diesem Sommer bin ich in der Firma verantwortlich für die
Lehrlingsausbildung. Ich durfte bereits von Januar bis August einen Praktikanten
betreuen und so erste Erfahrungen auf diesem Gebiet sammeln. Die Aufgabe,
jemandem das Programmieren beizubringen, der noch keine oder nur sehr wenig
Erfahrungen auf diesem Gebiet hat, ist sehr interessant ‒ und sehr fordernd.
Hier experimentiere ich mit verschiedenen Ansätzen, wobei ich nach fast einem
halben Jahr mehr offene Fragen als Antworten habe.

Im Frühling hat sich dann per Zufall eine Nebenbeschäftigung als
Berufsschullehrer ergeben: An einer Informationsveranstaltung für die neue
Bildungsverordnung der Informatiklehre wurde ich von meinem früheren
Klassenlehrer in der Berufsschule, der jetzt dort als Fachbereichsleiter in der
Informatik-Abteilung tätig ist, auf offene Lehrerstellen hingewiesen. Da hier
auch 20%-Pensen angeboten worden sind, und beruflich nur zu 80% tätig war,
ergriff ich diese Chance sofort.

So unterrichte ich seit Ende August ein Modul mit der Bezeichnung «Software mit
agilen Methoden entwickeln». Ich habe drei (recht unterschiedliche) Klassen zu
je zwei Wochenlektionen. Diese Aufgabe hat meine zweite Jahreshälfte stark
geprägt. Dies einerseits, weil ich einen grossen Teil meiner Freizeit für die
Vor- und Nachbereitung des Unterrichts verwende; andererseits, weil der Inhalt
dieses Moduls mich beschäftigte:

- Die Berufsschule ist Microsoft-Territorium. So bin ich nicht darum herum
  gekommen, mich seit 2005 das erste mal wieder mit C# zu beschäftigen.
  Besonders interessant finde ich die Programmiersprache nicht; ich sehe sie
  eher als eine Art _Microsoft Java_. Positiv überrascht hat mich jedoch an den
  aktuellen Versionen von .NET (5 und 6), dass man damit mittlerweile sehr gut
  unabhängig von Betriebssystem und IDE entwickeln kann. So darf C# in meinem
  Programmiersprachenrucksack gerne einen kleinen Platz einnehmen.
- Für die Unterrichtsgestaltung war für mich Git das A und O. Ich baute einen
  privaten Gitea-Server auf, dessen API mir bei der Administration von drei
  Klassen mit insgesamt 52 Schülern sehr stark behilflich war. Ich konnte einige
  Wissenlücken zum Thema Git schliessen und habe einige Unterlagen dazu
  erarbeitet, welche ich gelegentlich noch veröffentlichen möchte.
- Damit ich die gleichen Themen nicht immer wieder aufs neue erklären musste,
  habe ich einige Screencasts aufgenommen und diese auf einem eigens dafür
  geschaffenen
  [YouTube-Kanal](https://www.youtube.com/channel/UCPq4iLFbolH2deHLMVKha1A)
  veröffentlicht. Mit [OBS Studio](https://obsproject.com/) geht das mit wenig
  Aufwand, wobei ich auf einen Videoschnitt und besondere Effekte verzichte. In
  der Regel gelingen mir die Videos im ersten Versuch in einer einigermassen
  akzeptablen Qualität. Die Audioqualität meines Headsets war mir jedoch
  eindeutig zu schlecht, zumal meine Aussprache nicht sehr deutlich ist und ich
  mich beim Sprechen oft etwas überschlage. So habe ich mir den [Rode
  Podcaster](https://www.rode.com/microphones/podcaster) gegönnt, womit man doch
  sehr gute Ergebnisse erzielt. Ich trenne die Tonspur jeweils mit
  [`ffmpeg`](https://ffmpeg.org/) vom Video, importiere sie in
  [Audacity](https://www.audacityteam.org/) und mache sie mit dem
  Kompressor-Filter lauter. Diese Tonspur führe ich dann wieder mit dem Video
  zusammen, wiederum mit `ffmpeg`. Das ist die einzige Nachbereitung, die ich an
  meinen Videos vornehme.
- Um mich einmal etwas gründlicher mit dem Thema «Agile» (ich hasse die
  Substantivierung dieses Begriffs) zu befassen, habe ich [Clean
  Agile](https://www.informit.com/store/clean-agile-back-to-basics-9780135781869)
  von Robert C. Martin gelesen. (Der Autor geht mir mit seiner Boomer-Attitüde ‒
  *«If I achieve a test coverage of 95% in my hobby project, then you must be
  able to do so in the projects of your your stressful day-to-day job…»* ‒
  zwar ziemlich auf die Nerven, hat aber einen recht pragmatischen und
  vernünftigen Zugang zum Thema und predigt nicht einen einzelnen Ansatz als den
  einzig Wahren.) Hierzu habe ich auch eine sehr ausführliche
  [Zusammenfassung](https://github.com/patrickbucher/docs/blob/master/clean-agile/clean-agile.pdf)
  geschrieben und für den Unterricht noch auf Deutsch
  [übersetzt](https://github.com/patrickbucher/docs/blob/master/clean-agile/clean-agile-de.pdf).
  Mit der Zusammenfassung bin ich sehr zufrieden, mit der Übersetzung etwas
  weniger; doch sie erfüllt ihre Aufgabe, wobei mich die Lesefaulheit einiger
  Schüler doch etwas überrascht hat.

# Pen and Paper

Beim Zusammenfassen von _Clean Agile_ wollte ich sorgfältig vorgehen und habe
darum wieder einmal mit Tinte und Papier gearbeitet. So habe ich jeweils ein
Kapitel zuerst einmal komplett durchgelesen, und dann die einzelnen Unterkapitel
handschrifltich auf wenige Seiten zusammengefasst. Diese habe ich dann am
Computer eingetippt und daraus ein PDF generiert.

Ich habe dieses Vorgehen gewählt, da ich bei der Lektüre von Gerald Weinbergs
[How Software Is Built](https://leanpub.com/howsoftwareisbuilt/c/99Qmmhd5KHxg)
schlechte Erfahrungen machen musste: So habe ich das eBook am Computer gelesen
und gleich _on the fly_ zusammengefasst. Leider ist dabei eher ein
[Notizkonvolut](https://github.com/patrickbucher/docs/blob/master/weinberg/how-software-is-built.md)
als eine gut lesbare Zusammenfassung von Weinbergs Gedanken entstanden.

So habe ich mich entschieden, in Zukunft öfters den entschleunigten Ansatz mit
Tinte und Papier zu verfolgen. Hierzu habe ich mir auch zwei eher günstige
Füllfederhalter gekauft, mit denen ich einigermassen gerne schrieb, jedoch nicht
komplett zufrieden war. Durch das Lesen einiger von Dijkstras handschrifltich
verfasster [EWDs](https://www.cs.utexas.edu/users/EWD/) konnte ich mich dann
doch zur Anschaffung eines _Montblanc Meisterstück 149_ durchringen: bereut habe
ich es nicht, obwohl ich leider in der zweiten Jahreshälfte bedeutend weniger
zum Schreiben gekommen bin, als ich das gewollt hätte.

# Parerga und Paralipomena

Endlich habe ich mich im Frühling/Sommer einmal gründlich mit TLS befasst: Das
eBook [TLS Mastery](https://www.tiltedwindmillpress.com/product/tls/), das ich
auch sponserte, lag endlich bereit. Ich habe es sehr gründlich durchgearbeitet
und eine
[Zusammenfassung](https://github.com/patrickbucher/docs/blob/master/tls/tls-mastery.pdf)
mit weiteren Beispielen und Anhängen verfasst. Geschrieben habe ich sie wiederum
auf Papier, wobei ich beim Eintippen dann ebensoviel Arbeit durch das
Ausprobieren der Beispiele hatte. Jetzt fühle ich mich nicht nur mit TLS
einigermassen sattelfest, sondern verstehe auch ACME (Let's Encrypt) und habe
zum ersten mal einen kleinen DNS-Server aufgesetzt.

Leider war dann für mich in diesem Jahr was Systemadministration betrifft so
ziemlich die Luft raus. Mit [FreeBSD](https://www.freebsd.org/) und
[OpenBSD](https://www.openbsd.org/) habe ich mich dieses Jahr kaum beschäftigt.
Zwar leistet mein Datei- und Backupserver basierend auf billigster Hardware und
FreeBSD mit ZFS gute Dienste, daran angepasst habe ich jedoch wenig. Immerhin
laufen jetzt darauf verschiedenste Backup-Scripts, unter anderem
[back-my-git-up](https://github.com/patrickbucher/back-my-git-up), wozu ich die
[Suitable-Library](https://github.com/seantis/suitable) meines Arbeitgebers
verwenden konnte.

Mit [Radicale](https://radicale.org/v3.html) habe ich der Firma einen
CalDAV-Server zur Verfügung gestellt. Das ging so einfach, dass ich jetzt eine
private Instanz davon betreibe. So habe ich zum ersten mal in meinem Leben meine
Termine zentral abgelegt, ja sogar mit dem Smartphone synchronisiert. Auch
dieser Kalender wird von meinem Backup-Server gesichert, und zwar per `rsync`.

Somit habe ich eine produktive FreeBSD-Installation im Einsatz, die ich dieses
Jahr auf den Release 13.0 aktualisiert habe. OpenBSD habe ich leider derzeit gar
nicht mehr im Einsatz, was ich demnächst wieder ändern möchte. Habe ich im
Frühling noch auf OpenBSD Elixir gelernt, läuft bei mir mittlerweile wieder
überall Linux und Windows, letzteres wegen meiner Tätigkeit als
Berufsschullehrer. Ich werde meinen über fünfjährigen Desktop wohl Anfang
nächstes Jahr durch ein bereits angeschafftes Gerät von Lenovo ersetzen, da die
USB-Anschlüsse mittlerweile zu viele Probleme machen. Die Hardware kann ich aber
noch sehr gut als OpenBSD-Heimserver verwenden, denn CPU, Memory und die SSD
sind für Experimente durchaus noch gut genug.

Überhaupt sehne ich mich wieder etwas mehr nach Askese in der Informatik: Neben
OpenBSD drängt sich bei mir auch Go wieder etwas in den Vordergrund, zumal die
Einführung von Generics eine erneute Beschäftigung mit der Sprache für mich
nötig macht.  Ausserdem möchte ich Go vielleicht in der Berufsschule für das
kommende Cloud-Modul verwenden, das ich wohl rund um die Idee der [Twelve-Factor
App](https://12factor.net/) konzipieren werde. Da sich
[meillo](http://marmaro.de/) neustens auch mit Go befasst, ergibt sich hier ein
recht interessanten Austausch für mich.

Für eine systematische Auseinandersetzung mit [Puppet](https://puppet.com/)
konnte ich mich dieses Jahr leider nicht motivieren. Auch Themen wie Container
und die Cloud sind bei mir dieses Jahr liegengeblieben. Die geplante
LPIC-Zertifizierung ist an meiner oben bereits erwähnten Müdigkeit im
Zusammenhang mit Videokursen dieses Jahr gescheitert.

Auch für eine gründliche Auseinandersetzung
[SQLAlchemy](https://www.sqlalchemy.org/) konnte ich mich dieses Jahr nicht
motivieren. Dennoch komme ich mittlerweile einigermassen damit zurecht und kann
auch immer mehr explizite Queries durch die Nutzung der ORM-Features loswerden.
Hierdurch verstärkt sich meine ambivalente Haltung gegenüber ORMs: Zum ersten
mal siehe ich ihren positiven Nutzen; dennoch graut mir vor der ganzen
Komplexität, die dadurch erzeugt und damit versteckt wird.

Ansonsten habe ich mich noch kurz mit PHP und JavaScript befasst, dies in erster
Linie für meine Unterrichtstätigkeit im nächsten Frühling, wo ich ein
Praxismodul zum Thema Web-Entwicklung betreuen werde. Mit PHP habe ich innerlich
abgeschlossen und werde wohl eher die Idee von Full-Stack-JavaScript, d.h. nicht
nur im Browser, sondern auch serverseitig mit [Node.js](https://nodejs.org/en/)
oder [Deno](https://deno.land/) verfolgen.

Beinahe hätte ich es vergessen: 2021 habe ich endlich einmal
[plan9](https://9p.io/plan9/) ausprobiert, und zwar auf dem [Raspberry Pi
400](https://www.raspberrypi.com/products/raspberry-pi-400/). Viel ist nicht
daraus geworden, aber die Erfahrung möchte ich nicht missen: Diese Kombination
von Hardware und Betriebssystem versprüht etwas vom Pioniergeist der 80er-Jahre,
den ich nicht miterleben konnte.

In diesem Zusammenhang fallen mir noch einige Bücher ein, die ich in der zweiten
Jahreshälfte gelesen habe:
[Hackers](https://www.oreilly.com/library/view/hackers/9781449390259/), [Free as
in Freedom](https://www.oreilly.com/openbook/freedom/) und [Masters of
Doom](https://doomwiki.org/wiki/Masters_of_Doom). Diese Werke haben gemeinsam,
dass Durchbrüche und Höchstleistungen im Bereich der Programmierung der
körperlichen und sozialen Verwahrlosung dieser Hackertypen gegenübergestellt
werden. Diese Texte waren so für mich Motivation und Warnung zugleich. Mehrmals
wurde in diesem Kontext [Joseph
Weizenbaum](https://en.wikipedia.org/wiki/Joseph_Weizenbaum) zitiert, mit dessen
Werk ich mich auch erneut befassen möchte.

# Ausblick

Beruflich möchte ich mich 2022 einmal etwas gründlicher mit
[Ansible](https://www.ansible.com/) auseinandersetzen. Was Python betrifft,
sollte ich mich nicht nur mit der Sprache, sondern auch mit dem Ökosystem etwas
genauer befassen, z.B. mit dem Packaging und mit PyTest.

In der Berufsschule werde ich mich im ersten Halbjahr v.a. mit Web-Entwicklung,
d.h. mit JavaScript oder TypeScript auseinandersetzen müssen. An
[Express](https://expressjs.com/) habe ich recht gute Erinnerung. Bei der
funktionalen Programmierug mit JavaScript werde ich mich wohl etwas zurückhalten
müssen, sind doch Berufsschüler recht stark von der objektorientierten
Programmierung geprägt. Im zweiten Halbjahr soll dann das Cloud-Modul
stattfinden, wofür mir Go und [Heroku](https://www.heroku.com/) vorschwebt.

Privat möchte ich mich einmal gründlich mit Elixir befassen. Für mich ist das
die aussichtsreichste Plattform, um später einmal damit produktive Anwendungen
entwickeln zu können. Hier interessieren mich nicht nur die technischen Ideen ‒
funktionale und verteilte Programmierung, das Actor-Modell ‒ sondern auch
Frameworks wie [Phoenix](https://www.phoenixframework.org/) und
[Ecto](https://hexdocs.pm/ecto/Ecto.html). Ich bin gespannt, ob diese beiden
Technologien meine Vorurteile gegen Web-Frameworks und ORMs widerlegen können.
Wer weiss; vielleicht zündet die Beschäftigung damit bei mir endlich den Funken
um später meine vagen Pläne für die Selbständigkeit befeuern zu können.

SICP werde ich wohl auch 2022 nicht durcharbeiten können. Gerne hätte ich mich
einmal mit Knuths Werken (_The Art of Computer Programming_, _Concrete
Mathematics_) auseinandergesetzt; doch das muss wohl noch ein paar Jahre auf
meinem Büchergestellt reifen.

Mit Fremdsprachen sieht es auch im nächstn Jahr eher düster aus. Reisen dürfte
wohl immer noch zu umständlich oder zu schlecht planbar sein. Aber an ein
französisches oder russisches Buch könnte ich mich wieder einmal heranwagen.

Überhaupt möchte ich meine Bildschirmzeit etwas besser beschränken, zumal mich
immer häufiger Spannungskopfschmerzen plagen. Den Kindle habe ich immerhin
wieder durch gedruckte Bücher ersetzt. Das ist oftmals weniger komfortabel, aber
doch zielführender.

Auch sollte ich mich nächstes Jahr mehr mit didaktischen und pädagogischen
Fragen befassen, obwohl mir hierzu keine konkreten Themen oder Werke einfallen.
(Neben mir liegt Ivan Illichs [Entschulung der
Gesellschaft](https://www.perlentaucher.de/buch/ivan-illich/entschulung-der-gesellschaft.html),
das mich eher vom Thema weg- als darin einführt.)

An Büchern hat sich mittlerweile hier soviel angesammelt, dass ich mich eher mit
dem Verschenken und Entsorgen als mit dem Anschaffen beschäftigen sollte. Doch
auch dieses Problem werde ich 2022 nicht lösen können.

Sollte sich ab Frühling 2022 nicht mehr alle paar Tage ändern, was man unter
welchen Umständen noch tun darf, möchte ich wieder einmal den Ausstieg aus dem
News-Hamsterrad wagen; idealerweise für immer, damit ich endlich die Ruhe und
Zeit habe, um wieder einmal gründlich lesen zu können. Ja: _Konzentration_ wäre
ein guter Vorsatz fürs neue Jahr, wieder einmal…

# Fazit

Ein anstrengendes, arbeitsintensives und lehrreiches Jahr neigt sich dem Ende
zu. Wie ich schon im Untertitel vermutete, war 2021 dann doch nicht so schlecht,
wie mich das meine derzeitige Laune glauben lassen macht. Von daher hat sich das
Herunterschreiben meiner diesjährigen Leistungen (und das Ausbleiben
ebensolcher) doch gelohnt.

Die Zombieapokalypse ist ausgeblieben, ich und die Leute in meinem Umfeld sind
grösstenteils bei guter Gesundheit und in den Supermärkten erhält man immer noch
geniessbare Nahrungsmittel zu einigermassen stabilen Preisen. Das reduzierte
gesellschaftliche Leben lässt mir Zeit zum Lesen, Lernen und Spazieren. Ich
sollte diese Zeit nutzen; das ist das Beste, was ich tun kann.

# Nachtrag (21.12.2021)

Langsam lichtet sich der Schatten; einige Sachen sind bei meiner Aufstellung
vergessen gegangen:

- Ich habe mich kurz mit der Programmiersprache [Io](https://iolanguage.org/)
  befasst, um in der Berufsschule eine Lücke in der Jahresplanung schliessen zu
  können. Immerhin hat sich ein Schüler über die neue Perspektive auf OOP
  gefreut ‒ und darüber, dass schon sehr grosse Geister Kritik an OOP geübt
  haben. Ein paar Notizen und Übungen dazu kann ich gelegentlich
  veröffentlichen.
- Für die Bewertung einer kleinen Projektarbeit komme ich kaum um Excel herum.
  Ich möchte aber den Schülern ein Dokument mit der Bewertung abgeben können. So
  habe ich mir mit Python, OpenPyXL,
  [Jinja-Templates](https://jinja.palletsprojects.com/en/3.0.x/) und
  [Pandoc](https://pandoc.org/) ein kleines Werkzeug namens
  [gradedocs](https://github.com/patrickbucher/gradedocs) gebaut, das ich gerade
  erfolgreich einsetzen konnte.
- Nach der Lektüre von _Hackers_ habe ich endlich mal das [Game of
  Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in JavaScript
  ausprogrammiert. Man kann es unter
  [paeedubucher.ch/gol](https://paedubucher.ch/gol/) ausprobieren
  ([Code](https://github.com/patrickbucher/js-game-of-life) auf GitHub).
- Vor Jahren habe ich einmal bei Nassim Taleb gelesen, dass man Pi berechnen
  könne, indem man Pfeile auf ein Quadrat werfe, und dann schaue, welche
  innerhalb eines davon eingelassenen Kreises mit maximalem Durchmesser landen.
  Nun habe ich das einmal selber ausprogrammiert
  ([paedubucher.ch/simpi](https://paedubucher.ch/simpi/),
  [Code](https://github.com/patrickbucher/simulate-pi) auf GitHub). Es scheint
  tatsächlich einigermassen zu funktionieren.
- Ich wollte mich einmal umschauen, ob es für Go auch ein paar vernünftige
  Bibliotheken zur PDF-Generierung gibt. Ich bin dann auf
  [gopdf](https://github.com/signintech/gopdf) gestossen und habe daraus einen
  kleinen Flash-Card-Generator namens
  [verzettler](https://github.com/patrickbucher/verzettler) geschrieben. Das
  ganze ist ein Proof-of-Concept; mit langen Strings funktioniert es (noch)
  nicht; aber 2022 muss ich ja auch noch etwas zu tun haben…
- Obwohl es mir länger vorkommt, war es diesen Februar, als ich ein paar
  Routinen zum extrahieren von Daten aus dem DOM-Tree zu einer Go-Library
  verpackt habe: [htmlsqueeze](https://github.com/patrickbucher/htmlsqueeze).
  Ich setze das persönlich für ein Projekt ein. Die Weiterentwicklung könnte
  jedoch recht kompliziert werden, sofern ich noch weitere CSS-Selektoren
  (korrekt) umsetzen möchte.

Unproduktiv war das Jahr also nicht.
